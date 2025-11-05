package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/order/internal/model"

	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *order.OrderStatus) (*order.Empty, error) {
	if result := l.svcCtx.Db.Model(&model.OrderInfo{}).Where(&model.OrderInfo{
		OrderSn: in.OrderSn,
	}).Update("status", in.Status); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	}
	return &order.Empty{}, nil
}
