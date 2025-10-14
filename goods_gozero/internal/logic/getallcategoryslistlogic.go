package logic

import (
	"context"
	"encoding/json"
	"shop/goods_gozero/internal/model"

	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCategorysListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCategorysListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCategorysListLogic {
	return &GetAllCategorysListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品分类
func (l *GetAllCategorysListLogic) GetAllCategorysList(in *goods.Empty) (*goods.CategoryListResponse, error) {
	// todo: add your logic here and delete this line
	var categorys []model.Category
	//通过反向查询，查一级内目，查二级内目，查三级内目
	l.svcCtx.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)

	//json序列化
	b, _ := json.Marshal(&categorys)

	return &goods.CategoryListResponse{JsonData: string(b)}, nil
}
