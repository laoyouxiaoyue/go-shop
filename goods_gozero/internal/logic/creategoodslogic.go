package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsLogic {
	return &CreateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsLogic) CreateGoods(in *goods.CreateGoodsInfo) (*goods.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.GoodsInfoResponse{}, nil
}
