package logic

import (
	"context"
	"shop/order/internal/model"

	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemListLogic {
	return &CartItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 购物车
func (l *CartItemListLogic) CartItemList(in *order.UserInfo) (*order.CartItemListResponse, error) {
	var shopCars []model.ShoppingCart
	result := l.svcCtx.Db.Where(&model.ShoppingCart{User: in.Id}).Find(&shopCars)
	if result.Error != nil {
		return nil, result.Error
	}

	Rsp := order.CartItemListResponse{
		Total: int32(result.RowsAffected),
	}
	for _, value := range shopCars {
		Rsp.Data = append(Rsp.Data, &order.ShopCartInfoResponse{
			Id:      value.ID,
			UserId:  value.User,
			GoodsId: value.Goods,
			Nums:    value.Nums,
			Checked: value.Checked,
		})
	}
	return &Rsp, nil
}
