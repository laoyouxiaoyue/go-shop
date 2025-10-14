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

type DeleteCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryBrandLogic {
	return &DeleteCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryBrandLogic) DeleteCategoryBrand(in *goods.CategoryBrandRequest) (*goods.Empty, error) {
	if result := l.svcCtx.DB.Delete(&model.GoodsCategoryBrand{}, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	return &goods.Empty{}, nil
}
