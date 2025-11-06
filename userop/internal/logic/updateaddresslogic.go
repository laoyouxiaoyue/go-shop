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

type UpdateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAddressLogic {
	return &UpdateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAddressLogic) UpdateAddress(in *userop.GetAddressRequest) (*userop.Empty, error) {
	var address model.Address
	if result := l.svcCtx.Db.Where("id=? and user=?", in.Id, in.UserId).First(&address); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "收货地址记录不存在")
	}

	if in.Province != "" {
		address.Province = in.Province
	}
	if in.City != "" {
		address.City = in.City
	}
	if in.District != "" {
		address.District = in.District
	}
	if in.Address != "" {
		address.Address = in.Address
	}
	if in.SignerMobile != "" {
		address.SignerMobile = in.SignerMobile
	}
	if in.SignerName != "" {
		address.SignerName = in.SignerName
	}

	if err := l.svcCtx.Db.Save(&address).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "更新地址失败: %v", err)
	}

	return &userop.Empty{}, nil
}
