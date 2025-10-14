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

type CreateCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryBrandLogic {
	return &CreateCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryBrandLogic) CreateCategoryBrand(in *goods.CategoryBrandRequest) (*goods.CategoryBrandResponse, error) {
	var categroy model.Category

	//在创建品牌分类前，必须有对应商品分类,没有商品分类, 是不允许创建品牌分类的
	//查询商品分类
	if result := l.svcCtx.DB.First(&categroy, in.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	//查询品牌
	var brand model.Brands
	if result := l.svcCtx.DB.First(&brand, in.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categroyBrand := model.GoodsCategoryBrand{
		CategoryID: categroy.ID,
		BrandsID:   brand.ID,
	}

	//在对应表中插入记录
	l.svcCtx.DB.Save(&categroyBrand)

	return &goods.CategoryBrandResponse{
		Id: int32(categroyBrand.ID),
	}, nil
}
