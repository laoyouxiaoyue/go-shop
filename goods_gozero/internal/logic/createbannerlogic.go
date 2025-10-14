package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/model"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBannerLogic {
	return &CreateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBannerLogic) CreateBanner(in *goods.BannerRequest) (*goods.BannerResponse, error) {
	if result := l.svcCtx.DB.First(model.Banner{}, in.Id); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "轮播图已存在")
	}
	var Banner model.Banner
	Banner.ID = in.Id
	Banner.Index = in.Index
	Banner.Image = in.Image
	Banner.Url = in.Url
	if result := l.svcCtx.DB.Create(&Banner); result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &goods.BannerResponse{
		Id:    Banner.ID,
		Index: Banner.Index,
		Image: Banner.Image,
		Url:   Banner.Url,
	}, nil
}
