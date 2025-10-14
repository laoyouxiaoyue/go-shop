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

type CreateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBrandLogic {
	return &CreateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBrandLogic) CreateBrand(in *goods.BrandRequest) (*goods.BrandInfoResponse, error) {
	var brand *model.Brands
	// 使用事务
	err := l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 检查品牌名称是否已存在
		var existing model.Brands
		err := tx.Where("name = ?", in.Name).First(&existing).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			zap.S().Error("查询品牌是否存在失败", zap.Error(err), zap.String("name", in.Name))
			return status.Error(codes.Internal, "数据库查询失败")
		}

		if err == nil {
			zap.S().Warn("品牌名称已存在", zap.String("name", in.Name))
			return status.Error(codes.AlreadyExists, "品牌已存在")
		}

		// 2. 创建品牌
		brand = &model.Brands{
			Name: in.Name,
			Logo: in.Logo,
		}

		err = tx.Create(brand).Error
		if err != nil {
			zap.S().Error("创建品牌失败", zap.Error(err), zap.Any("brand", brand))
			return status.Error(codes.Internal, "数据库操作失败")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	zap.S().Info("创建品牌成功", zap.String("name", in.Name), zap.Uint("id", uint(brand.ID)))
	return &goods.BrandInfoResponse{
		Id:   brand.ID,
		Logo: brand.Logo,
		Name: brand.Name,
	}, nil
}
