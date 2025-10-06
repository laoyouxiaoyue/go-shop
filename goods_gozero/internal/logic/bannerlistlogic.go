package logic

import (
	"context"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerListLogic {
	return &BannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 轮播图
func (l *BannerListLogic) BannerList(in *goods.Empty) (*goods.BannerListResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.BannerListResponse{}, nil
}
