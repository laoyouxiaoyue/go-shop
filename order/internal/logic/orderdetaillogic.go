package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
	"shop/order/internal/model"

	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func Model2Rsp(goods model.OrderGoods) *order.OrderItemResponse {
	return &order.OrderItemResponse{
		Id:         goods.ID,
		OrderId:    goods.Order,
		GoodsName:  goods.GoodsName,
		GoodsPrice: goods.GoodsPrice,
		GoodsImage: goods.GoodsImage,
		Nums:       goods.Nums,
		GoodsId:    goods.Goods,
	}
}
func (l *OrderDetailLogic) OrderDetail(in *order.OrderRequest) (*order.OrderInfoDetailResponse, error) {
	var orderInfo model.OrderInfo
	var rsp order.OrderInfoDetailResponse
	if result := l.svcCtx.Db.Where(&model.OrderInfo{
		BaseModel: model.BaseModel{
			ID: in.Id,
		},
		User: in.UserId,
	}).First(&orderInfo); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "订单记录不存在")
	}
	OrderInfo := order.OrderInfoResponse{
		Id:      orderInfo.ID,
		UserId:  orderInfo.User,
		OrderSn: orderInfo.OrderSn,
		PayType: orderInfo.PayType,
		Status:  orderInfo.Status,
		Post:    orderInfo.Post,
		Total:   orderInfo.OrderMount,
		Address: orderInfo.Address,
		Name:    orderInfo.SignerName,
		Mobile:  orderInfo.SingerMobile,
		AddTime: fmt.Sprintf("%d-%d-%d %d:%d:%d", orderInfo.CreatedAt.Year(), orderInfo.CreatedAt.Month(), orderInfo.CreatedAt.Day()),
	}

	rsp.OrderInfo = &OrderInfo

	var GoodsList []model.OrderGoods
	if result := l.svcCtx.Db.Where(
		&model.OrderGoods{
			Order: in.Id,
		}).First(&OrderInfo); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	for _, goods := range GoodsList {
		rsp.Goods = append(rsp.Goods, Model2Rsp(goods))
	}
	return &rsp, nil
}
