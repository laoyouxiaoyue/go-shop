package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetGoodsLogic {
	return &BatchGetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (l *BatchGetGoodsLogic) BatchGetGoods(in *goods.BatchGoodsIdInfo) (*goods.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.GoodsListResponse{}, nil
}
