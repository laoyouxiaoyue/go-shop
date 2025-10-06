package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &goods.SubCategoryListResponse{}, nil
}
