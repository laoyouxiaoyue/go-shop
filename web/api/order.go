package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
	"shop/order/orderclient"
	"shop/web/forms"
	"shop/web/utils"
)

// 创建order gRPC客户端连接
func createOrderClient() (orderclient.Order, error) {
	addr, err := utils.DiscoverAddr("order")
	if err != nil {
		return nil, err
	}
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{addr},
		NonBlock:  true,
		Timeout:   int64(time.Second * 3),
	})
	return orderclient.NewOrder(cli), nil
}

// 购物车列表
func GetCartList(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		zap.S().Errorw("[GetCartList] 连接order服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	uidStr := c.DefaultQuery("userId", "0")
	uid, _ := strconv.Atoi(uidStr)
	rsp, err := cli.CartItemList(context.Background(), &orderclient.UserInfo{Id: int32(uid)})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 加入购物车
func CreateCartItem(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.CartItemForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	rsp, err := cli.CreateCarItem(context.Background(), &orderclient.CartItemRequest{
		Id:         form.Id,
		UserId:     form.UserId,
		GoodsId:    form.GoodsId,
		GoodsName:  form.GoodsName,
		GoodsImage: form.GoodsImage,
		GoodsPrice: form.GoodsPrice,
		Nums:       form.Nums,
		Checked:    form.Checked,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 更新购物车
func UpdateCartItem(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.CartItemForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	_, err = cli.UpdateCartItem(context.Background(), &orderclient.CartItemRequest{
		Id:         form.Id,
		UserId:     form.UserId,
		GoodsId:    form.GoodsId,
		GoodsName:  form.GoodsName,
		GoodsImage: form.GoodsImage,
		GoodsPrice: form.GoodsPrice,
		Nums:       form.Nums,
		Checked:    form.Checked,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 删除购物车项
func DeleteCartItem(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.CartItemForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	_, err = cli.DeleteCartItem(context.Background(), &orderclient.CartItemRequest{Id: form.Id, UserId: form.UserId, GoodsId: form.GoodsId})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 创建订单
func CreateOrder(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.CreateOrderForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	rsp, err := cli.CreateOrder(context.Background(), &orderclient.OrderRequest{
		Id:      form.Id,
		UserId:  form.UserId,
		Address: form.Address,
		Name:    form.Name,
		Mobile:  form.Mobile,
		Post:    form.Post,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 订单列表
func GetOrderList(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.OrderFilterForm
	// 允许 query 或 json，两者择一
	if c.ContentType() == "application/json" {
		_ = c.ShouldBindJSON(&form)
	} else {
		_ = c.ShouldBindQuery(&form)
	}
	if form.Pages == 0 {
		form.Pages = 1
	}
	if form.PagePerNums == 0 {
		form.PagePerNums = 10
	}
	rsp, err := cli.OrderList(context.Background(), &orderclient.OrderFilterRequest{
		UserId:      form.UserId,
		Pages:       form.Pages,
		PagePerNums: form.PagePerNums,
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 订单详情
func GetOrderDetail(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	rsp, err := cli.OrderDetail(context.Background(), &orderclient.OrderRequest{Id: int32(id)})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

// 更新订单状态
func UpdateOrderStatus(c *gin.Context) {
	cli, err := createOrderClient()
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var form forms.UpdateOrderStatusForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	_, err = cli.UpdateOrderStatus(context.Background(), &orderclient.OrderStatus{Id: form.Id, Status: form.Status})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
