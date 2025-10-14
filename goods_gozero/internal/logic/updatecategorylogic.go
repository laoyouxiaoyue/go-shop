package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *goods.CategoryInfoRequest) (*goods.Empty, error) {
	var category model.Category
	if result := l.svcCtx.DB.First(&category, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	if in.Name != "" {
		category.Name = in.Name
	}

	if in.ParentCategory != 0 {
		//类目级别
		category.ParentCategoryID = in.ParentCategory
	}
	if in.Level != 0 {
		category.Level = in.Level
	}
	if in.IsTab {
		category.IsTab = in.IsTab
	}

	l.svcCtx.DB.Save(&category)

	return &goods.Empty{}, nil
}
