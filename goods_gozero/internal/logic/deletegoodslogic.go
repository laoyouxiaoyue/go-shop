package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsLogic) DeleteGoods(in *goods.DeleteGoodsInfo) (*goods.Empty, error) {
	// todo: add your logic here and delete this line

	return &goods.Empty{}, nil
}
