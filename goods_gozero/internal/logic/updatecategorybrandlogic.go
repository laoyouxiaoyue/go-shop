package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &goods.Empty{}, nil
}
