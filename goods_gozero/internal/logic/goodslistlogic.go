package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品接口
func (l *GoodsListLogic) GoodsList(in *goods.GoodsFilterRequest) (*goods.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.GoodsListResponse{}, nil
}
