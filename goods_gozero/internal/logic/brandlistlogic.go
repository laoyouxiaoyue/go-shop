package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌
func (l *BrandListLogic) BrandList(in *goods.BrandFilterRequest) (*goods.BrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.BrandListResponse{}, nil
}
