package logic

import (
	"context"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *goods.CategoryInfoRequest) (*goods.CategoryInfoResponse, error) {
	category := model.Category{}
	category.Name = in.Name
	category.Level = in.Level
	if in.Level != 1 {
		//不是一级类目，需要将指向一级类目
		category.ParentCategoryID = in.ParentCategory

		//category.ParentCategoryID = req.ParentCategory
	}
	category.IsTab = in.IsTab

	//global.DB.Save(&category)
	l.svcCtx.DB.Model(model.Category{}).Save(&category)

	return &goods.CategoryInfoResponse{Id: int32(category.ID)}, nil
}
