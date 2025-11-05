package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	goods2 "shop/goods_gozero/goods"
	"shop/inventory/inventory"
	"shop/order/internal/model"
	"time"

	"shop/order/internal/svc"
	"shop/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func GenerationSn(userId int32) string {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	//年月日时分秒 + 用户id + 两位随机数
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), userId, rand.Intn(90)+10)
	return orderSn
}

// 订单
func (l *CreateOrderLogic) CreateOrder(in *order.OrderRequest) (*order.OrderInfoResponse, error) {

	var goodsIds []int32
	var shopCarts []model.ShoppingCart

	if result := l.svcCtx.Db.Where(&model.ShoppingCart{
		User:    in.UserId,
		Checked: true,
	}).Find(&shopCarts); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "没有选中的商品")
	}
	goodsNumsMap := make(map[int32]int32)
	for _, shopCart := range shopCarts {
		goodsIds = append(goodsIds, shopCart.Goods)
		goodsNumsMap[shopCart.Goods] += shopCart.Nums
	}

	goods, err := l.svcCtx.Goods.BatchGetGoods(context.Background(), &goods2.BatchGoodsIdInfo{
		Id: goodsIds,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "批量查询商品信息失败")
	}
	var total float32 = 0
	var orderGoods []*model.OrderGoods
	var goodsInvInfo []*inventory.GoodsInventoryInfo
	for _, good := range goods.Data {
		total += good.ShopPrice * float32(goodsNumsMap[good.Id])
		orderGoods = append(orderGoods, &model.OrderGoods{
			Goods:      good.Id,
			GoodsName:  good.Name,
			GoodsImage: good.GoodsFrontImage,
			GoodsPrice: good.ShopPrice,
			Nums:       goodsNumsMap[good.Id],
		})
		goodsInvInfo = append(goodsInvInfo, &inventory.GoodsInventoryInfo{
			GoodsId: good.Id,
			Num:     goodsNumsMap[good.Id],
		})
	}

	if _, err = l.svcCtx.Inv.Sell(context.Background(), &inventory.SellInfo{
		GoodsInfo: goodsInvInfo,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "扣减库存失败")
	}
	tx := l.svcCtx.Db.Begin()
	orders := model.OrderInfo{
		OrderSn:      GenerationSn(in.UserId),
		OrderMount:   total,
		Address:      in.Address,
		SignerName:   in.Name,
		SingerMobile: in.Mobile,
		Post:         in.Post,
	}

	if result := tx.Save(&orders); result.RowsAffected == 0 {
		tx.Rollback()
	}

	for _, good := range orderGoods {
		good.Order = orders.ID
	}

	if result := tx.CreateInBatches(orderGoods, 100); result.RowsAffected == 0 {
		tx.Rollback()
	}

	if result := tx.Where(&model.ShoppingCart{
		User:    in.UserId,
		Checked: true,
	}).Delete(&model.ShoppingCart{}); result.RowsAffected == 0 {
		tx.Rollback()
	}
	tx.Commit()
	return &order.OrderInfoResponse{
		Id:      orders.ID,
		OrderSn: orders.OrderSn,
		Total:   orders.OrderMount,
	}, nil
}
