package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/userop/internal/model"

	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFavDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFavDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFavDetailLogic {
	return &GetUserFavDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFavDetailLogic) GetUserFavDetail(in *userop.UserFavRequest) (*userop.Empty, error) {
	var userfav model.UserFav
	if result := l.svcCtx.Db.Where("goods=? and user=?", in.GoodsId, in.UserId).First(&userfav); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &userop.Empty{}, nil
}
