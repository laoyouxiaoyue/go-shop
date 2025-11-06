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

type CreateAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAddressLogic {
	return &CreateAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAddressLogic) CreateAddress(in *userop.GetAddressRequest) (*userop.AddressResponse, error) {
	var address model.Address
	address.User = in.UserId
	address.Province = in.Province
	address.City = in.City
	address.District = in.District
	address.Address = in.Address
	address.SignerName = in.SignerName
	address.SignerMobile = in.SignerMobile

	if err := l.svcCtx.Db.Save(&address).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "新建地址失败: %v", err)
	}
	return &userop.AddressResponse{Id: address.ID}, nil
}
