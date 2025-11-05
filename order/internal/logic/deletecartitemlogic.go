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

type DeleteCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartItemLogic {
	return &DeleteCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCartItemLogic) DeleteCartItem(in *order.CartItemRequest) (*order.Empty, error) {
	if result := l.svcCtx.Db.Where(model.ShoppingCart{
		User:  in.UserId,
		Goods: in.GoodsId,
	}).Delete(&model.ShoppingCart{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	return &order.Empty{}, nil
}
