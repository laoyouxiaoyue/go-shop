package logic

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBrandLogic) DeleteBrand(in *goods.BrandRequest) (*goods.Empty, error) {
	err := l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		// 检查品牌是否存在并删除
		result := tx.Where("id = ?", in.Id).Delete(&model.Brands{})
		if result.Error != nil {
			zap.S().Error("删除品牌失败", zap.Error(result.Error), zap.Uint("id", uint(in.Id)))
			return status.Error(codes.Internal, "数据库操作失败")
		}

		// 检查是否删除了记录
		if result.RowsAffected == 0 {
			zap.S().Warn("品牌不存在", zap.Uint("id", uint(in.Id)))
			return status.Error(codes.NotFound, "品牌不存在")
		}

		zap.S().Info("删除品牌成功", zap.Uint("id", uint(in.Id)))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &goods.Empty{}, nil
}
