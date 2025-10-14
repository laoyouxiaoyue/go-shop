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

type UpdateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandLogic {
	return &UpdateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandLogic) UpdateBrand(in *goods.BrandRequest) (*goods.Empty, error) {
	err := l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查品牌是否存在
		var existingBrand model.Brands
		err := tx.Where("id = ?", in.Id).First(&existingBrand).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				zap.S().Warn("品牌不存在", zap.Uint("id", uint(in.Id)))
				return status.Error(codes.NotFound, "品牌不存在")
			}
			zap.S().Error("查询品牌失败", zap.Error(err), zap.Uint("id", uint(in.Id)))
			return status.Error(codes.Internal, "数据库查询失败")
		}

		// 2. 如果名称有变化，检查新名称是否已存在
		if in.Name != existingBrand.Name {
			var count int64
			err := tx.Model(&model.Brands{}).
				Where("name = ? AND id != ?", in.Name, in.Id).
				Count(&count).Error
			if err != nil {
				zap.S().Error("检查品牌名称冲突失败", zap.Error(err))
				return status.Error(codes.Internal, "数据库查询失败")
			}
			if count > 0 {
				zap.S().Warn("品牌名称已存在", zap.String("name", in.Name))
				return status.Error(codes.AlreadyExists, "品牌名称已存在")
			}
		}

		// 3. 更新品牌信息
		updates := map[string]interface{}{
			"name": in.Name,
			"logo": in.Logo,
		}

		result := tx.Model(&model.Brands{}).Where("id = ?", in.Id).Updates(updates)
		if result.Error != nil {
			zap.S().Error("更新品牌失败", zap.Error(result.Error), zap.Uint("id", uint(in.Id)))
			return status.Error(codes.Internal, "数据库操作失败")
		}

		zap.S().Info("更新品牌成功", zap.Uint("id", uint(in.Id)))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &goods.Empty{}, nil
}
