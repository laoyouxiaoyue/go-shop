package logic

import (
	"context"
	"shop/order/internal/model"
	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCarItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCarItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCarItemLogic {
	return &CreateCarItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCarItemLogic) CreateCarItem(in *order.CartItemRequest) (*order.ShopCartInfoResponse, error) {
	var shopCart model.ShoppingCart
	if result := l.svcCtx.Db.Where(&model.ShoppingCart{
		Goods: in.GoodsId,
		User:  in.UserId,
	}).First(&shopCart); result.RowsAffected == 1 {
		shopCart.Nums += in.Nums
	} else {
		shopCart.Nums = in.Nums
		shopCart.Goods = in.GoodsId
		shopCart.User = in.UserId
		shopCart.Checked = false
	}
	l.svcCtx.Db.Save(&shopCart)
	return &order.ShopCartInfoResponse{
		Id: shopCart.ID,
	}, nil
}
