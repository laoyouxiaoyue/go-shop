package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(in *goods.CreateGoodsInfo) (*goods.Empty, error) {
	// todo: add your logic here and delete this line

	return &goods.Empty{}, nil
}
