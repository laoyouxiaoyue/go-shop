package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
	"net/http"
	"shop/goods_gozero/goodsclient"
	"shop/web/forms"
	"shop/web/utils"
	"strconv"
	"time"
)

// 创建goods gRPC客户端连接
func createGoodsClient() (goodsclient.Goods, error) {
	addr, err := utils.DiscoverAddr("goods")
	if err != nil {
		return nil, err
	}
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{addr},
		NonBlock:  true,
		Timeout:   int64(time.Second * 3),
	})
	return goodsclient.NewGoods(cli), nil
}

// 商品列表
func GetGoodsList(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetGoodsList] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	// 获取查询参数
	priceMin, _ := strconv.Atoi(c.DefaultQuery("priceMin", "0"))
	priceMax, _ := strconv.Atoi(c.DefaultQuery("priceMax", "0"))
	isHot := c.DefaultQuery("isHot", "false") == "true"
	isNew := c.DefaultQuery("isNew", "false") == "true"
	isTab := c.DefaultQuery("isTab", "false") == "true"
	topCategory, _ := strconv.Atoi(c.DefaultQuery("topCategory", "0"))
	pages, _ := strconv.Atoi(c.DefaultQuery("pages", "1"))
	pagePerNums, _ := strconv.Atoi(c.DefaultQuery("pagePerNums", "10"))
	keyWords := c.DefaultQuery("keyWords", "")
	brand, _ := strconv.Atoi(c.DefaultQuery("brand", "0"))

	req := &goodsclient.GoodsFilterRequest{
		PriceMin:    int32(priceMin),
		PriceMax:    int32(priceMax),
		IsHot:       isHot,
		IsNew:       isNew,
		IsTab:       isTab,
		TopCategory: int32(topCategory),
		Pages:       int32(pages),
		PagePerNums: int32(pagePerNums),
		KeyWords:    keyWords,
		Brand:       int32(brand),
	}

	rsp, err := goodsClient.GoodsList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetGoodsList] 获取商品列表错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  rsp.Data,
	})
}

// 批量获取商品信息
func BatchGetGoods(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[BatchGetGoods] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var ids []int32
	if err := c.ShouldBindJSON(&ids); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.BatchGoodsIdInfo{
		Id: ids,
	}

	rsp, err := goodsClient.BatchGetGoods(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[BatchGetGoods] 批量获取商品错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  rsp.Data,
	})
}

// 获取商品详情
func GetGoodsDetail(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetGoodsDetail] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	req := &goodsclient.GoodInfoRequest{
		Id: int32(id),
	}

	rsp, err := goodsClient.GetGoodsDetail(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetGoodsDetail] 获取商品详情错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, rsp)
}

// 获取所有分类
func GetAllCategorysList(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetAllCategorysList] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	req := &goodsclient.Empty{}

	rsp, err := goodsClient.GetAllCategorysList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetAllCategorysList] 获取分类列表错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":    rsp.Total,
		"data":     rsp.Data,
		"jsonData": rsp.JsonData,
	})
}

// 获取子分类
func GetSubCategory(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetSubCategory] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Query("id"))
	level, _ := strconv.Atoi(c.DefaultQuery("level", "1"))

	req := &goodsclient.CategoryListRequest{
		Id:    int32(id),
		Level: int32(level),
	}

	rsp, err := goodsClient.GetSubCategory(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetSubCategory] 获取子分类错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":        rsp.Total,
		"info":         rsp.Info,
		"subCategorys": rsp.SubCategorys,
	})
}

// 获取品牌列表
func GetBrandList(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetBrandList] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	pages, _ := strconv.Atoi(c.DefaultQuery("pages", "1"))
	pagePerNums, _ := strconv.Atoi(c.DefaultQuery("pagePerNums", "10"))

	req := &goodsclient.BrandFilterRequest{
		Pages:       int32(pages),
		PagePerNums: int32(pagePerNums),
	}

	rsp, err := goodsClient.BrandList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetBrandList] 获取品牌列表错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  rsp.Data,
	})
}

// 获取轮播图列表
func GetBannerList(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetBannerList] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	req := &goodsclient.Empty{}

	rsp, err := goodsClient.BannerList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetBannerList] 获取轮播图列表错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  rsp.Data,
	})
}

// 根据分类获取品牌
func GetCategoryBrandList(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[GetCategoryBrandList] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.CategoryInfoRequest{
		Id: int32(id),
	}

	rsp, err := goodsClient.GetCategoryBrandList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[GetCategoryBrandList] 获取分类品牌列表错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": rsp.Total,
		"data":  rsp.Data,
	})
}

// ==================== 商品管理CRUD接口 ====================

// 创建商品
func CreateGoods(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[CreateGoods] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var createForm forms.CreateGoodsForm
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CreateGoodsInfo{
		Name:            createForm.Name,
		GoodsSn:         createForm.GoodsSn,
		Stocks:          createForm.Stocks,
		MarketPrice:     createForm.MarketPrice,
		ShopPrice:       createForm.ShopPrice,
		GoodsBrief:      createForm.GoodsBrief,
		GoodsDesc:       createForm.GoodsDesc,
		ShipFree:        createForm.ShipFree,
		Images:          createForm.Images,
		DescImages:      createForm.DescImages,
		GoodsFrontImage: createForm.GoodsFrontImage,
		IsNew:           createForm.IsNew,
		IsHot:           createForm.IsHot,
		OnSale:          createForm.OnSale,
		CategoryId:      createForm.CategoryId,
		BrandId:         createForm.BrandId,
	}

	rsp, err := goodsClient.CreateGoods(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[CreateGoods] 创建商品错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusCreated, rsp)
}

// 更新商品
func UpdateGoods(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[UpdateGoods] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var updateForm forms.UpdateGoodsForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CreateGoodsInfo{
		Id:              int32(id),
		Name:            updateForm.Name,
		GoodsSn:         updateForm.GoodsSn,
		Stocks:          updateForm.Stocks,
		MarketPrice:     updateForm.MarketPrice,
		ShopPrice:       updateForm.ShopPrice,
		GoodsBrief:      updateForm.GoodsBrief,
		GoodsDesc:       updateForm.GoodsDesc,
		ShipFree:        updateForm.ShipFree,
		Images:          updateForm.Images,
		DescImages:      updateForm.DescImages,
		GoodsFrontImage: updateForm.GoodsFrontImage,
		IsNew:           updateForm.IsNew,
		IsHot:           updateForm.IsHot,
		OnSale:          updateForm.OnSale,
		CategoryId:      updateForm.CategoryId,
		BrandId:         updateForm.BrandId,
	}

	_, err = goodsClient.UpdateGoods(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[UpdateGoods] 更新商品错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "商品更新成功",
	})
}

// 删除商品
func DeleteGoods(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[DeleteGoods] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.DeleteGoodsInfo{
		Id: int32(id),
	}

	_, err = goodsClient.DeleteGoods(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[DeleteGoods] 删除商品错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "商品删除成功",
	})
}

// ==================== 分类管理CRUD接口 ====================

// 创建分类
func CreateCategory(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[CreateCategory] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var createForm forms.CreateCategoryForm
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CategoryInfoRequest{
		Name:           createForm.Name,
		ParentCategory: createForm.ParentCategory,
		Level:          createForm.Level,
		IsTab:          createForm.IsTab,
	}

	rsp, err := goodsClient.CreateCategory(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[CreateCategory] 创建分类错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusCreated, rsp)
}

// 更新分类
func UpdateCategory(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[UpdateCategory] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var updateForm forms.UpdateCategoryForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CategoryInfoRequest{
		Id:             int32(id),
		Name:           updateForm.Name,
		ParentCategory: updateForm.ParentCategory,
		Level:          updateForm.Level,
		IsTab:          updateForm.IsTab,
	}

	_, err = goodsClient.UpdateCategory(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[UpdateCategory] 更新分类错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "分类更新成功",
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[DeleteCategory] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.DeleteCategoryRequest{
		Id: int32(id),
	}

	_, err = goodsClient.DeleteCategory(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[DeleteCategory] 删除分类错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "分类删除成功",
	})
}

// ==================== 品牌管理CRUD接口 ====================

// 创建品牌
func CreateBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[CreateBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var createForm forms.CreateBrandForm
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.BrandRequest{
		Name: createForm.Name,
		Logo: createForm.Logo,
	}

	rsp, err := goodsClient.CreateBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[CreateBrand] 创建品牌错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusCreated, rsp)
}

// 更新品牌
func UpdateBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[UpdateBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var updateForm forms.UpdateBrandForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.BrandRequest{
		Id:   int32(id),
		Name: updateForm.Name,
		Logo: updateForm.Logo,
	}

	_, err = goodsClient.UpdateBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[UpdateBrand] 更新品牌错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "品牌更新成功",
	})
}

// 删除品牌
func DeleteBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[DeleteBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.BrandRequest{
		Id: int32(id),
	}

	_, err = goodsClient.DeleteBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[DeleteBrand] 删除品牌错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "品牌删除成功",
	})
}

// ==================== 轮播图管理CRUD接口 ====================

// 创建轮播图
func CreateBanner(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[CreateBanner] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var createForm forms.CreateBannerForm
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.BannerRequest{
		Index: createForm.Index,
		Image: createForm.Image,
		Url:   createForm.Url,
	}

	rsp, err := goodsClient.CreateBanner(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[CreateBanner] 创建轮播图错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusCreated, rsp)
}

// 更新轮播图
func UpdateBanner(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[UpdateBanner] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var updateForm forms.UpdateBannerForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.BannerRequest{
		Id:    int32(id),
		Index: updateForm.Index,
		Image: updateForm.Image,
		Url:   updateForm.Url,
	}

	_, err = goodsClient.UpdateBanner(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[UpdateBanner] 更新轮播图错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "轮播图更新成功",
	})
}

// 删除轮播图
func DeleteBanner(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[DeleteBanner] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.BannerRequest{
		Id: int32(id),
	}

	_, err = goodsClient.DeleteBanner(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[DeleteBanner] 删除轮播图错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "轮播图删除成功",
	})
}

// ==================== 分类品牌关联管理CRUD接口 ====================

// 创建分类品牌关联
func CreateCategoryBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[CreateCategoryBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	var createForm forms.CreateCategoryBrandForm
	if err := c.ShouldBindJSON(&createForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CategoryBrandRequest{
		CategoryId: createForm.CategoryId,
		BrandId:    createForm.BrandId,
	}

	rsp, err := goodsClient.CreateCategoryBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[CreateCategoryBrand] 创建分类品牌关联错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusCreated, rsp)
}

// 更新分类品牌关联
func UpdateCategoryBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[UpdateCategoryBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var updateForm forms.UpdateCategoryBrandForm
	if err := c.ShouldBindJSON(&updateForm); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}

	req := &goodsclient.CategoryBrandRequest{
		Id:         int32(id),
		CategoryId: updateForm.CategoryId,
		BrandId:    updateForm.BrandId,
	}

	_, err = goodsClient.UpdateCategoryBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[UpdateCategoryBrand] 更新分类品牌关联错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "分类品牌关联更新成功",
	})
}

// 删除分类品牌关联
func DeleteCategoryBrand(c *gin.Context) {
	goodsClient, err := createGoodsClient()
	if err != nil {
		zap.S().Errorw("[DeleteCategoryBrand] 连接goods服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	req := &goodsclient.CategoryBrandRequest{
		Id: int32(id),
	}

	_, err = goodsClient.DeleteCategoryBrand(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[DeleteCategoryBrand] 删除分类品牌关联错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "分类品牌关联删除成功",
	})
}
