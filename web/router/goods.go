package router

import (
	"github.com/gin-gonic/gin"
	"shop/web/api"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		// ==================== 商品管理接口 ====================
		GoodsRouter.GET("list", api.GetGoodsList)         // 获取商品列表
		GoodsRouter.POST("batch", api.BatchGetGoods)      // 批量获取商品信息
		GoodsRouter.GET("detail/:id", api.GetGoodsDetail) // 获取商品详情
		GoodsRouter.POST("", api.CreateGoods)             // 创建商品
		GoodsRouter.PUT(":id", api.UpdateGoods)           // 更新商品
		GoodsRouter.DELETE(":id", api.DeleteGoods)        // 删除商品

		// ==================== 分类管理接口 ====================
		GoodsRouter.GET("categories", api.GetAllCategorysList)   // 获取所有分类
		GoodsRouter.GET("subcategories", api.GetSubCategory)     // 获取子分类
		GoodsRouter.POST("categories", api.CreateCategory)       // 创建分类
		GoodsRouter.PUT("categories/:id", api.UpdateCategory)    // 更新分类
		GoodsRouter.DELETE("categories/:id", api.DeleteCategory) // 删除分类

		// ==================== 品牌管理接口 ====================
		GoodsRouter.GET("brands", api.GetBrandList)       // 获取品牌列表
		GoodsRouter.POST("brands", api.CreateBrand)       // 创建品牌
		GoodsRouter.PUT("brands/:id", api.UpdateBrand)    // 更新品牌
		GoodsRouter.DELETE("brands/:id", api.DeleteBrand) // 删除品牌

		// ==================== 轮播图管理接口 ====================
		GoodsRouter.GET("banners", api.GetBannerList)       // 获取轮播图列表
		GoodsRouter.POST("banners", api.CreateBanner)       // 创建轮播图
		GoodsRouter.PUT("banners/:id", api.UpdateBanner)    // 更新轮播图
		GoodsRouter.DELETE("banners/:id", api.DeleteBanner) // 删除轮播图

		// ==================== 分类品牌关联管理接口 ====================
		GoodsRouter.GET("category/:id/brands", api.GetCategoryBrandList)  // 根据分类获取品牌
		GoodsRouter.POST("category-brand", api.CreateCategoryBrand)       // 创建分类品牌关联
		GoodsRouter.PUT("category-brand/:id", api.UpdateCategoryBrand)    // 更新分类品牌关联
		GoodsRouter.DELETE("category-brand/:id", api.DeleteCategoryBrand) // 删除分类品牌关联
	}
}
