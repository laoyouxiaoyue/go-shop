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

type UpdateCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartItemLogic) UpdateCartItem(in *order.CartItemRequest) (*order.Empty, error) {
	var shopCart model.ShoppingCart
	if result := l.svcCtx.Db.Where(&model.ShoppingCart{
		Goods: in.GoodsId,
		User:  in.UserId,
	}).First(&shopCart); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车不存在该记录")
	}

	shopCart.Checked = in.Checked
	if in.Nums > 0 {
		shopCart.Nums = in.Nums
	}
	l.svcCtx.Db.Save(&shopCart)
	return &order.Empty{}, nil
}
