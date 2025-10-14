package model

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

// Category 分类栏
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20) comment '分类名称';not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`                                                        //表外键
	ParentCategory   *Category   `json:"-"`                                                             //指向的父级分类栏
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"` //填装子分类数据
	Level            int32       `gorm:"type:int comment '分类级别';not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"` //是否上架
}

// Brands 品牌信息
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20) comment '品牌名称';not null"`
	Logo string `gorm:"type:varchar(200) comment '品牌logo图片方式展示';default:'';not null"`
}

// GoodsCategoryBrand 商品和品牌的关系,需要建立关系，使用联合index
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

// TableName 默认名称
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// Banner 轮播图展示
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200) comment '展示图片';not null"`
	Url   string `gorm:"type:varchar(200) comment '商品链接';not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

// Goods 商品信息
type Goods struct {
	BaseModel
	//外键
	CategoryID int32 `gorm:"type:int;not null"` //商品分类外键
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"` //品牌分类外键
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"` //是否上架
	ShipFree bool `gorm:"default:false;not null"` //是否包邮
	IsNew    bool `gorm:"default:false;not null"` //是否新品
	IsHot    bool `gorm:"default:false;not null"` //是否热销

	Name            string   `gorm:"type:varchar(50) comment '商品名称';not null"`       //商品名称
	GoodsSn         string   `gorm:"type:varchar(50) comment '商家内部商品编号';not null"`   //商家内部商品编号
	ClickNum        int32    `gorm:"type:int comment '商品点击量';default:0;not null"`    //商品点击量
	SoldNum         int32    `gorm:"type:int comment '商品销量';default:0;not null"`     //商品销量
	FavNum          int32    `gorm:"type:int comment '商品收藏量';default:0;not null"`    //商品收藏量
	MarketPrice     float32  `gorm:"not null"`                                       //标记
	ShopPrice       float32  `gorm:"not null"`                                       //实际价
	GoodsBrief      string   `gorm:"type:varchar(100) comment '商品介绍';not null"`      //商品介绍
	Images          GormList `gorm:"type:varchar(1000) comment '商品图片链接';not null"`   //商品图片链接
	DescImages      GormList `gorm:"type:varchar(1000) comment '商品详细图片链接';not null"` //商品详细图片链接
	GoodsFrontImage string   `gorm:"type:varchar(200) comment '商品封面链接';not null"`    //商品封面链接
}

type GormList []string

func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), g)
}
func (g *GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}
