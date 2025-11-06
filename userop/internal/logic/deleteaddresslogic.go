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

type DeleteAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressLogic {
	return &DeleteAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAddressLogic) DeleteAddress(in *userop.GetAddressRequest) (*userop.Empty, error) {
	result := l.svcCtx.Db.Where("id=? and user=?", in.Id, in.UserId).Delete(&model.Address{})
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "删除地址失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收货地址记录不存在")
	}
	return &userop.Empty{}, nil
}
