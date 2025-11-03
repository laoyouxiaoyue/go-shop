package logic

import (
	"context"
	"shop/inventory/internal/model"

	"shop/inventory/internal/svc"
	"shop/inventory/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetInvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInvLogic {
	return &SetInvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetInvLogic) SetInv(in *inventory.GoodsInventoryInfo) (*inventory.Empty, error) {
	var inventory1 model.Inventory
	l.svcCtx.Db.Where(&model.Inventory{Goods: in.GoodsId}).First(&inventory1)
	inventory1.Goods = in.GoodsId
	inventory1.Stocks = in.Num
	l.svcCtx.Db.Save(&inventory1)
	return &inventory.Empty{}, nil
}
