package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsLogic) DeleteGoods(in *goods.DeleteGoodsInfo) (*goods.Empty, error) {
	tx := l.svcCtx.DB.Begin()
	if result := tx.Delete(&model.Goods{BaseModel: model.BaseModel{ID: in.Id}}); result.RowsAffected == 0 {
		if result.Error != nil {
			tx.Rollback()
			return nil, result.Error
		}
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	tx.Commit()
	return &goods.Empty{}, nil
}
