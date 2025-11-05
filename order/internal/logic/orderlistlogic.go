package logic

import (
	"context"
	"shop/order/internal/model"

	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func OrderInfo2Rsp(o model.OrderInfo) *order.OrderInfoResponse {
	return &order.OrderInfoResponse{
		Id:      o.ID,
		UserId:  o.User,
		OrderSn: o.OrderSn,
		PayType: o.PayType,
		Status:  o.Status,
		Post:    o.Post,
		Address: o.Address,
		Name:    o.SignerName,
		Mobile:  o.SingerMobile,
		Total:   o.OrderMount,
		AddTime: o.CreatedAt.Format("2006-04-02 15:04:05"),
	}
}
func (l *OrderListLogic) OrderList(in *order.OrderFilterRequest) (*order.OrderListResponse, error) {
	var orders []model.OrderInfo
	var rsp order.OrderListResponse
	var total int64
	l.svcCtx.Db.Where(&model.OrderInfo{
		User: in.UserId,
	}).Find(&model.OrderInfo{}).Count(&total)
	rsp.Total = int32(total)
	for _, order1 := range orders {
		rsp.Data = append(rsp.Data, OrderInfo2Rsp(order1))
	}
	return &rsp, nil
}
