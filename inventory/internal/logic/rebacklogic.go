package logic

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/inventory/internal/model"

	"shop/inventory/internal/svc"
	"shop/inventory/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type RebackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRebackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RebackLogic {
	return &RebackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RebackLogic) Reback(in *inventory.SellInfo) (*inventory.Empty, error) {
	tx := l.svcCtx.Db.Begin()
	for _, goodsInfo := range in.GoodsInfo {
		var inventory model.Inventory
		for {
			if result := l.svcCtx.Db.Where(&model.Inventory{Goods: goodsInfo.GoodsId}).First(&inventory); result.RowsAffected == 0 {
				//失败进行事务回滚
				tx.Rollback()
				return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
			}
			inventory.Stocks += goodsInfo.Num
			if result := tx.Model(&model.Inventory{}).Select("Stocks", "Version").Where("goods = ? and version = ?",
				goodsInfo.GoodsId, inventory.Version).Updates(model.Inventory{Stocks: inventory.Stocks, Version: inventory.Version + 1}); result.RowsAffected == 0 {
				zap.S().Info("库存归还失败")
			} else {
				break
			}
		}
	}
	//提交事务
	tx.Commit()
	return &inventory.Empty{}, nil
}
