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

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *goods.DeleteCategoryRequest) (*goods.Empty, error) {
	// todo: add your logic here and delete this line
	if result := l.svcCtx.DB.Model(model.Category{}).Delete(&model.Category{}, in.Id); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &goods.Empty{}, nil
}
