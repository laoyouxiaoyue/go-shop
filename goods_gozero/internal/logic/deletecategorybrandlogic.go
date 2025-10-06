package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryBrandLogic {
	return &DeleteCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryBrandLogic) DeleteCategoryBrand(in *goods.CategoryBrandRequest) (*goods.Empty, error) {
	// todo: add your logic here and delete this line

	return &goods.Empty{}, nil
}
