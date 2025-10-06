package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryBrandLogic {
	return &CreateCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryBrandLogic) CreateCategoryBrand(in *goods.CategoryBrandRequest) (*goods.CategoryBrandResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.CategoryBrandResponse{}, nil
}
