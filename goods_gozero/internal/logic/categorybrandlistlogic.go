package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryBrandListLogic {
	return &CategoryBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌分类
func (l *CategoryBrandListLogic) CategoryBrandList(in *goods.CategoryBrandFilterRequest) (*goods.CategoryBrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.CategoryBrandListResponse{}, nil
}
