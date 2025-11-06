package logic

import (
	"context"
	"shop/userop/internal/model"

	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressListLogic {
	return &GetAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressListLogic) GetAddressList(in *userop.GetAddressRequest) (*userop.AddressListResponse, error) {
	var addressList []model.Address
	var AddressListResponse userop.AddressListResponse
	if result := l.svcCtx.Db.Where(&model.Address{User: in.UserId}).Find(&addressList); result.RowsAffected != 0 {
		AddressListResponse.Total = int32(result.RowsAffected)
	}

	for _, address := range addressList {
		AddressListResponse.Data = append(AddressListResponse.Data, &userop.AddressResponse{
			Id:           address.ID,
			UserId:       address.User,
			Province:     address.Province,
			City:         address.City,
			District:     address.District,
			Address:      address.Address,
			SignerMobile: address.SignerMobile,
			SignerName:   address.SignerName,
		})
	}
	return &AddressListResponse, nil
}
