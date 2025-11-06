package logic

import (
	"context"
	"shop/userop/internal/model"

	"shop/userop/internal/svc"
	"shop/userop/userop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavListLogic {
	return &GetFavListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavListLogic) GetFavList(in *userop.UserFavRequest) (*userop.UserFavListResponse, error) {
	var userFavs []model.UserFav
	l.svcCtx.Db.Where(&model.UserFav{
		User:  in.UserId,
		Goods: in.GoodsId,
	}).Find(&userFavs)
	var userfavrsp []*userop.UserFavResponse
	for i := 0; i < len(userFavs); i++ {
		userfavrsp = append(
			userfavrsp,
			&userop.UserFavResponse{
				UserId:  userFavs[i].User,
				GoodsId: userFavs[i].Goods,
			})
	}
	return &userop.UserFavListResponse{
		Data: userfavrsp,
	}, nil
}
