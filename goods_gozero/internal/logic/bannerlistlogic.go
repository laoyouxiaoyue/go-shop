package logic

import (
	"context"
	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/model"
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
	bannerListResponse := goods.BannerListResponse{}

	var banners []model.Banner
	result := l.svcCtx.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*goods.BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &goods.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil
}
