package model

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32  `gorm:"type:int;not null" json:"parent_category_id"`
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool  `gorm:"not null;default:false" json:"isTab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Logo string `gorm:"type:varchar(255);not null;default:''" json:"logo"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique;" json:"category_id"`
	Category   Category

	BrandID int32 `gorm:"type:int;index:idx_category_brand,unique;" json:"brand_id"`
	Brand   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null" json:"image"`
	Url   string `gorm:"type:varchar(255);not null" json:"url"`
	Index int32  `gorm:"type:int;not null;default:1" json:"index"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null;" json:"category_id"`
	Category   Category

	BrandID int32 `gorm:"type:int;not null;" json:"brand_id"`
	Brand   Brands

	OnSale   bool `gorm:"default:false;not null" json:"on_sale"`
	ShipFree bool `gorm:"default:false;not null" json:"ship_free"`
	IsNew    bool `gorm:"default:false;not null" json:"is_new"`
	IsHot    bool `gorm:"default:false;not null" json:"is_hot"`

	Name            string   `gorm:"type:varchar(50);not null" json:"name"`
	GoodsSn         string   `gorm:"type:varchar(50);not null" json:"goods_sn"`
	ClickNum        int32    `gorm:"type:int;not null;default:0" json:"click_num"`
	SoldNum         int32    `gorm:"type:int;not null;default:0" json:"sold_num"`
	FavNum          int32    `gorm:"type:int;not null;default:0" json:"fav_num"`
	MarketPrice     float32  `gorm:"type:float;not null" json:"market_price"`
	ShopPrice       float32  `gorm:"type:float;not null" json:"shop_price"`
	GoodsBrief      string   `gorm:"type:varchar(255);not null" json:"goods_brief"`
	Images          GormList `gorm:"type:varchar(1000);not null" json:"images"`
	DescImages      GormList `gorm:"type:varchar(1000);not null" json:"descimages"`
	GoodsFrontImage string   `gorm:"type:varchar(255);not null" json:"goods_front_image"`
}
