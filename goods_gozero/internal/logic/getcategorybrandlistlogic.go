package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryBrandListLogic {
	return &GetCategoryBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过category获取brands
func (l *GetCategoryBrandListLogic) GetCategoryBrandList(in *goods.CategoryInfoRequest) (*goods.BrandListResponse, error) {
	return nil, nil

	//var categoryBrands []model.GoodsCategoryBrand
	//var categoryBrandListResponse goods.CategoryBrandListResponse
	//
	//result := l.svcCtx.DB.Find(&categoryBrands)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	//categoryBrandListResponse.Total = int32(result.RowsAffected)
	//
	////分页
	//l.svcCtx.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)
	//
	//var CategoryBrandResponse []*goods.CategoryBrandResponse
	//for _, categoryBrand := range categoryBrands {
	//	CategoryBrandResponse = append(CategoryBrandResponse, &goods.CategoryBrandResponse{
	//		Id: categoryBrand.CategoryID,
	//		Brand: &goods.BrandInfoResponse{
	//			Id:   categoryBrand.Brands.ID,
	//			Name: categoryBrand.Brands.Name,
	//			Logo: categoryBrand.Brands.Logo,
	//		},
	//		Category: &goods.CategoryInfoResponse{
	//			Id:             categoryBrand.Category.ID,
	//			Name:           categoryBrand.Category.Name,
	//			Level:          categoryBrand.Category.Level,
	//			IsTab:          categoryBrand.Category.IsTab,
	//			ParentCategory: categoryBrand.Category.ParentCategoryID,
	//		},
	//	})
	//}
	//categoryBrandListResponse.Data = CategoryBrandResponse
	//
	//return &categoryBrandListResponse, nil
}
