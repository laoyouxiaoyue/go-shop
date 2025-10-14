package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsLogic {
	return &CreateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsLogic) CreateGoods(in *goods.CreateGoodsInfo) (*goods.GoodsInfoResponse, error) {
	var category model.Category
	if result := l.svcCtx.DB.First(&category, in.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := l.svcCtx.DB.First(&brand, in.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	var good model.Goods

	good.Brands = brand
	good.BrandsID = brand.ID
	good.Category = category
	good.CategoryID = category.ID

	good.IsNew = in.IsNew
	good.IsHot = in.IsHot
	good.OnSale = in.OnSale
	good.ShipFree = in.ShipFree

	good.ID = in.Id
	good.Name = in.Name
	good.GoodsSn = in.GoodsSn
	good.MarketPrice = in.MarketPrice
	good.ShopPrice = in.ShopPrice
	good.GoodsBrief = in.GoodsBrief
	good.GoodsFrontImage = in.GoodsFrontImage
	good.DescImages = in.DescImages
	good.Images = in.Images

	//这里需要利用钩子和事务将mysql和es数据保持一致性
	tx := l.svcCtx.DB.Begin()
	result := tx.Save(&good)
	if result.Error != nil {
		//业务回滚
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()
	return &goods.GoodsInfoResponse{
		Id: good.ID,
	}, nil
}
