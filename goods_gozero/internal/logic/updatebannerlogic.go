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

type UpdateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBannerLogic {
	return &UpdateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBannerLogic) UpdateBanner(in *goods.BannerRequest) (*goods.Empty, error) {
	var banner model.Banner
	if result := l.svcCtx.DB.First(&banner, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	if in.Image != "" {
		banner.Image = in.Image
	}
	if in.Index >= 0 {
		banner.Index = in.Index
	}
	if in.Url != "" {
		banner.Url = in.Url
	}
	l.svcCtx.DB.Save(&banner)
	return &goods.Empty{}, nil
}
