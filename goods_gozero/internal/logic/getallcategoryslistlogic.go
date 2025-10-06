package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCategorysListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCategorysListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCategorysListLogic {
	return &GetAllCategorysListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品分类
func (l *GetAllCategorysListLogic) GetAllCategorysList(in *goods.Empty) (*goods.CategoryListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.CategoryListResponse{}, nil
}
