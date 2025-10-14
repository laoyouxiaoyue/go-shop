package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品接口
func (l *GoodsListLogic) GoodsList(in *goods.GoodsFilterRequest) (*goods.GoodsListResponse, error) {
	db := l.svcCtx.DB.Model(&model.Goods{})
	if in.KeyWords != "" {
		db = db.Where("name LIKE ?", "%"+in.KeyWords+"%")
	}
	if in.IsTab == true {
		db = db.Where("is_tab = true")
	}
	if in.IsHot == true {
		db = db.Where("is_hot = true")
	}
	if in.IsNew == true {
		db = db.Where("is_new = true")
	}
	if in.PriceMin > 0 {
		db = db.Where("shop_price >= ?", in.PriceMin)
	}
	if in.PriceMax > 0 {
		db = db.Where("shop_price <= ?", in.PriceMax)
	}
	if in.Brand > 0 {
		db = db.Where("brand_id = ?", in.Brand)
	}

	if in.TopCategory > 0 {
		var category model.Category
		if result := l.svcCtx.DB.First(&category, in.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}
		var SubQuery string
		//根据类目级别返回查询子分类
		if category.Level == 1 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id in (SELECT id FROM category WHERE parent_category_id = %d)", in.TopCategory)
		} else if category.Level == 2 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE parent_category_id = %d", in.TopCategory)
		} else if category.Level == 3 {
			SubQuery = fmt.Sprintf("SELECT id FROM category WHERE id = %d", in.TopCategory)
		}
		db = db.Where(SubQuery)
	}
	var res []*model.Goods
	goodsListResponse := &goods.GoodsListResponse{}
	db.Find(&res)
	goodsListResponse.Total = int32(len(res))
	for _, good := range res {
		GoodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &GoodsInfoResponse)
	}

	return goodsListResponse, nil
}
func ModelToResponse(good *model.Goods) goods.GoodsInfoResponse {
	return goods.GoodsInfoResponse{
		Id:              good.ID,
		CategoryId:      good.CategoryID,
		Name:            good.Name,
		GoodsSn:         good.GoodsSn,
		ClickNum:        good.ClickNum,
		SoldNum:         good.SoldNum,
		FavNum:          good.FavNum,
		MarketPrice:     good.MarketPrice,
		ShopPrice:       good.ShopPrice,
		GoodsBrief:      good.GoodsBrief,
		ShipFree:        good.ShipFree,
		GoodsFrontImage: good.GoodsFrontImage,
		IsNew:           good.IsNew,
		IsHot:           good.IsHot,
		OnSale:          good.OnSale,
		DescImages:      good.DescImages,
		Images:          good.Images,
		Category: &goods.CategoryBriefInfoResponse{
			Id:   good.Category.ID,
			Name: good.Category.Name,
		},
		Brand: &goods.BrandInfoResponse{
			Id:   good.Brands.ID,
			Name: good.Brands.Name,
			Logo: good.Brands.Logo,
		},
	}
}
