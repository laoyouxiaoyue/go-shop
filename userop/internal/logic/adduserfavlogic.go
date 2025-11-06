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

type AddUserFavLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserFavLogic {
	return &AddUserFavLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserFavLogic) AddUserFav(in *userop.UserFavRequest) (*userop.Empty, error) {
	var userFav model.UserFav
	if result := l.svcCtx.Db.Where("user=? and goods=?", in.UserId, in.GoodsId).First(&userFav); result.Error == nil {
		return nil, status.Errorf(codes.AlreadyExists, "此收藏记录已存在")
	}

	if err := l.svcCtx.Db.Save(&model.UserFav{
		User:  in.UserId,
		Goods: in.GoodsId,
	}).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "创建收藏记录失败: %v", err)
	}
	return &userop.Empty{}, nil
}
