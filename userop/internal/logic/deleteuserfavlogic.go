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

type DeleteUserFavLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserFavLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFavLogic {
	return &DeleteUserFavLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserFavLogic) DeleteUserFav(in *userop.UserFavRequest) (*userop.Empty, error) {
	//Unscoped 设置了唯一索引使用硬删除
	result := l.svcCtx.Db.Unscoped().Where("goods=? and user=?", in.GoodsId, in.UserId).Delete(&model.UserFav{})
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "删除收藏失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &userop.Empty{}, nil
}
