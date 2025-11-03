package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/inventory/internal/model"

	"shop/inventory/internal/svc"
	"shop/inventory/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvDetailLogic {
	return &InvDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvDetailLogic) InvDetail(in *inventory.GoodsInventoryInfo) (*inventory.GoodsInventoryInfo, error) {
	var req model.Inventory
	if result := l.svcCtx.Db.Where(&model.Inventory{Goods: in.GoodsId}).First(&req); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "未找到对应的库存商品")
	}
	return &inventory.GoodsInventoryInfo{
		GoodsId: req.Goods,
		Num:     req.Stocks,
	}, nil
}
