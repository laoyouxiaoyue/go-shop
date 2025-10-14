package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"shop/goods_gozero/goods"
	"shop/goods_gozero/internal/model"
	"shop/goods_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自定义错误码
var (
	ErrInvalidParams = status.Error(codes.InvalidArgument, "请求参数不合法")
	ErrDBQuery       = status.Error(codes.Internal, "数据库查询失败")
	ErrDataNotFound  = status.Error(codes.NotFound, "数据不存在")
)

func (l *BrandListLogic) BrandList(in *goods.BrandFilterRequest) (*goods.BrandListResponse, error) {
	// 参数验证
	var BrandList goods.BrandListResponse
	var brands []model.Brands
	var BrandInfo []*goods.BrandInfoResponse

	result := l.svcCtx.DB.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	l.svcCtx.DB.Model(&brands).Count(&total)

	l.svcCtx.DB.Scopes(l.Paginate(int(in.Pages), int(in.PagePerNums))).Find(&brands)

	for _, brand := range brands {
		BrandInfo = append(BrandInfo, &goods.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	BrandList.Total = int32(total)
	BrandList.Data = BrandInfo
	return &BrandList, nil
}

func (l *BrandListLogic) Paginate(page, pageSize int) func(dao *gorm.DB) *gorm.DB {
	return func(dao *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return dao.Offset(offset).Limit(pageSize)
	}
}
