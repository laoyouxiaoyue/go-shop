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

type UpdateCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryBrandLogic {
	return &UpdateCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryBrandLogic) UpdateCategoryBrand(in *goods.CategoryBrandRequest) (*goods.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand
	if result := l.svcCtx.DB.First(&categoryBrand, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	//查询商品分类
	var category model.Category
	if result := l.svcCtx.DB.First(&category, in.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	//查询品牌
	var brand model.Brands
	if result := l.svcCtx.DB.First(&brand, in.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = in.CategoryId
	categoryBrand.BrandsID = in.BrandId

	l.svcCtx.DB.Save(&categoryBrand)

	return &goods.Empty{}, nil
}
