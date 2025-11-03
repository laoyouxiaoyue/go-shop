package logic

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/inventory/internal/model"

	"shop/inventory/internal/svc"
	"shop/inventory/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellLogic {
	return &SellLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SellLogic) Sell(in *inventory.SellInfo) (*inventory.Empty, error) {
	tx := l.svcCtx.Db.Begin()
	var mutexs []*redsync.Mutex
	for _, goodsInfo := range in.GoodsInfo {
		var inventory model.Inventory
		mutex := l.svcCtx.Lock.NewMutex(fmt.Sprintf("goods_%d", goodsInfo.GoodsId))

		if err := mutex.Lock(); err != nil {
			return nil, status.Errorf(codes.Internal, "获取redis分布式锁异常")
		}
		if result := l.svcCtx.Db.Where(&model.Inventory{Goods: goodsInfo.GoodsId}).First(&inventory); result.RowsAffected == 0 {
			//失败进行事务回滚
			tx.Rollback()
			return nil, status.Errorf(codes.InvalidArgument, "库存信息不存在")
		}
		if inventory.Stocks < goodsInfo.Num {
			//失败进行事务回滚
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		inventory.Stocks -= goodsInfo.Num
		tx.Save(&inventory)
		mutexs = append(mutexs, mutex)
	}
	//提交事
	tx.Commit()

	for _, mutex := range mutexs {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			return nil, status.Errorf(codes.Internal, "释放redis分布式锁异常")
		}
	}
	return &inventory.Empty{}, nil
}
