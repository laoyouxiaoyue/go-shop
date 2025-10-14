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

type GetSubCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCategoryLogic {
	return &GetSubCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取子分类
func (l *GetSubCategoryLogic) GetSubCategory(in *goods.CategoryListRequest) (*goods.SubCategoryListResponse, error) {
	categoryListResponse := goods.SubCategoryListResponse{}

	var category model.Category
	//查询分类是否存在
	//查询当前目录
	if result := l.svcCtx.DB.Find(&category, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}
	categoryListResponse.Info = &goods.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	//当前目录子目录
	var subCategorys []model.Category
	var categoryInfoResponse []*goods.CategoryInfoResponse
	perloads := "SubCategory"
	if category.Level == 1 {
		perloads = "SubCategory.SubCategory"
	}
	l.svcCtx.DB.Where(&model.Category{ParentCategoryID: in.Id}).Preload(perloads).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		categoryInfoResponse = append(categoryInfoResponse, &goods.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategorys = categoryInfoResponse
	return &categoryListResponse, nil
}
