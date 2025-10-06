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
	// todo: add your logic here and delete this line

	return &goods.BrandListResponse{}, nil
}
