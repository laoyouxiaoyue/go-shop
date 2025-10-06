package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBrandLogic) DeleteBrand(in *goods.BrandRequest) (*goods.Empty, error) {
	// todo: add your logic here and delete this line

	return &goods.Empty{}, nil
}
