package forms

// 商品筛选表单
type GoodsFilterForm struct {
	PriceMin    int    `form:"priceMin" json:"priceMin" binding:"min=0"`
	PriceMax    int    `form:"priceMax" json:"priceMax" binding:"min=0"`
	IsHot       bool   `form:"isHot" json:"isHot"`
	IsNew       bool   `form:"isNew" json:"isNew"`
	IsTab       bool   `form:"isTab" json:"isTab"`
	TopCategory int    `form:"topCategory" json:"topCategory" binding:"min=0"`
	Pages       int    `form:"pages" json:"pages" binding:"min=1"`
	PagePerNums int    `form:"pagePerNums" json:"pagePerNums" binding:"min=1,max=100"`
	KeyWords    string `form:"keyWords" json:"keyWords"`
	Brand       int    `form:"brand" json:"brand" binding:"min=0"`
}

// 批量获取商品ID表单
type BatchGoodsIdForm struct {
	Ids []int32 `json:"ids" binding:"required,min=1"`
}

// 分类查询表单
type CategoryQueryForm struct {
	Id    int `form:"id" json:"id" binding:"required,min=1"`
	Level int `form:"level" json:"level" binding:"min=1,max=3"`
}

// 品牌筛选表单
type BrandFilterForm struct {
	Pages       int `form:"pages" json:"pages" binding:"min=1"`
	PagePerNums int `form:"pagePerNums" json:"pagePerNums" binding:"min=1,max=100"`
}

// ==================== 商品管理表单 ====================

// 创建商品表单
type CreateGoodsForm struct {
	Name            string   `json:"name" binding:"required,min=1,max=100"`
	GoodsSn         string   `json:"goodsSn" binding:"required,min=1,max=50"`
	Stocks          int32    `json:"stocks" binding:"required,min=0"`
	MarketPrice     float32  `json:"marketPrice" binding:"required,min=0"`
	ShopPrice       float32  `json:"shopPrice" binding:"required,min=0"`
	GoodsBrief      string   `json:"goodsBrief" binding:"required,min=1,max=200"`
	GoodsDesc       string   `json:"goodsDesc" binding:"required,min=1"`
	ShipFree        bool     `json:"shipFree"`
	Images          []string `json:"images" binding:"required,min=1"`
	DescImages      []string `json:"descImages"`
	GoodsFrontImage string   `json:"goodsFrontImage" binding:"required,min=1"`
	IsNew           bool     `json:"isNew"`
	IsHot           bool     `json:"isHot"`
	OnSale          bool     `json:"onSale"`
	CategoryId      int32    `json:"categoryId" binding:"required,min=1"`
	BrandId         int32    `json:"brandId" binding:"required,min=1"`
}

// 更新商品表单
type UpdateGoodsForm struct {
	Name            string   `json:"name" binding:"required,min=1,max=100"`
	GoodsSn         string   `json:"goodsSn" binding:"required,min=1,max=50"`
	Stocks          int32    `json:"stocks" binding:"required,min=0"`
	MarketPrice     float32  `json:"marketPrice" binding:"required,min=0"`
	ShopPrice       float32  `json:"shopPrice" binding:"required,min=0"`
	GoodsBrief      string   `json:"goodsBrief" binding:"required,min=1,max=200"`
	GoodsDesc       string   `json:"goodsDesc" binding:"required,min=1"`
	ShipFree        bool     `json:"shipFree"`
	Images          []string `json:"images" binding:"required,min=1"`
	DescImages      []string `json:"descImages"`
	GoodsFrontImage string   `json:"goodsFrontImage" binding:"required,min=1"`
	IsNew           bool     `json:"isNew"`
	IsHot           bool     `json:"isHot"`
	OnSale          bool     `json:"onSale"`
	CategoryId      int32    `json:"categoryId" binding:"required,min=1"`
	BrandId         int32    `json:"brandId" binding:"required,min=1"`
}

// ==================== 分类管理表单 ====================

// 创建分类表单
type CreateCategoryForm struct {
	Name           string `json:"name" binding:"required,min=1,max=50"`
	ParentCategory int32  `json:"parentCategory" binding:"min=0"`
	Level          int32  `json:"level" binding:"required,min=1,max=3"`
	IsTab          bool   `json:"isTab"`
}

// 更新分类表单
type UpdateCategoryForm struct {
	Name           string `json:"name" binding:"required,min=1,max=50"`
	ParentCategory int32  `json:"parentCategory" binding:"min=0"`
	Level          int32  `json:"level" binding:"required,min=1,max=3"`
	IsTab          bool   `json:"isTab"`
}

// ==================== 品牌管理表单 ====================

// 创建品牌表单
type CreateBrandForm struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
	Logo string `json:"logo" binding:"required,min=1"`
}

// 更新品牌表单
type UpdateBrandForm struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
	Logo string `json:"logo" binding:"required,min=1"`
}

// ==================== 轮播图管理表单 ====================

// 创建轮播图表单
type CreateBannerForm struct {
	Index int32  `json:"index" binding:"required,min=0"`
	Image string `json:"image" binding:"required,min=1"`
	Url   string `json:"url" binding:"required,min=1"`
}

// 更新轮播图表单
type UpdateBannerForm struct {
	Index int32  `json:"index" binding:"required,min=0"`
	Image string `json:"image" binding:"required,min=1"`
	Url   string `json:"url" binding:"required,min=1"`
}

// ==================== 分类品牌关联管理表单 ====================

// 创建分类品牌关联表单
type CreateCategoryBrandForm struct {
	CategoryId int32 `json:"categoryId" binding:"required,min=1"`
	BrandId    int32 `json:"brandId" binding:"required,min=1"`
}

// 更新分类品牌关联表单
type UpdateCategoryBrandForm struct {
	CategoryId int32 `json:"categoryId" binding:"required,min=1"`
	BrandId    int32 `json:"brandId" binding:"required,min=1"`
}
