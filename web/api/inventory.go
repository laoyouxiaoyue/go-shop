package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
	"shop/inventory/inventoryclient"
	"shop/web/forms"
	"shop/web/utils"
)

// 创建inventory gRPC客户端连接
func createInventoryClient() (inventoryclient.Inventory, error) {
	addr, err := utils.DiscoverAddr("inventory")
	if err != nil {
		return nil, err
	}
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Endpoints: []string{addr},
		NonBlock:  true,
		Timeout:   int64(time.Second * 3),
	})
	return inventoryclient.NewInventory(cli), nil
}

// 设置库存
func SetInventory(c *gin.Context) {
	cli, err := createInventoryClient()
	if err != nil {
		zap.S().Errorw("[SetInventory] 连接inventory服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var req forms.SetInventoryForm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	_, err = cli.SetInv(context.Background(), &inventoryclient.GoodsInventoryInfo{GoodsId: req.GoodsId, Num: req.Num})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 查询库存详情
func GetInventoryDetail(c *gin.Context) {
	cli, err := createInventoryClient()
	if err != nil {
		zap.S().Errorw("[GetInventoryDetail] 连接inventory服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	idStr := c.Param("goodsId")
	gid, _ := strconv.Atoi(idStr)
	rsp, err := cli.InvDetail(context.Background(), &inventoryclient.GoodsInventoryInfo{GoodsId: int32(gid)})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"goodsId": rsp.GoodsId, "num": rsp.Num})
}

// 扣减库存（支持批量）
func SellInventory(c *gin.Context) {
	cli, err := createInventoryClient()
	if err != nil {
		zap.S().Errorw("[SellInventory] 连接inventory服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var req forms.InventoryBatchForm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	items := make([]*inventoryclient.GoodsInventoryInfo, 0, len(req.GoodsInfo))
	for _, it := range req.GoodsInfo {
		items = append(items, &inventoryclient.GoodsInventoryInfo{GoodsId: it.GoodsId, Num: it.Num})
	}
	_, err = cli.Sell(context.Background(), &inventoryclient.SellInfo{GoodsInfo: items, OrderSn: req.OrderSn})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

// 归还库存（支持批量）
func RebackInventory(c *gin.Context) {
	cli, err := createInventoryClient()
	if err != nil {
		zap.S().Errorw("[RebackInventory] 连接inventory服务错误", "err", err)
		HandleGrpcErrorToHttp(err, c)
		return
	}
	var req forms.InventoryBatchForm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
		return
	}
	items := make([]*inventoryclient.GoodsInventoryInfo, 0, len(req.GoodsInfo))
	for _, it := range req.GoodsInfo {
		items = append(items, &inventoryclient.GoodsInventoryInfo{GoodsId: it.GoodsId, Num: it.Num})
	}
	_, err = cli.Reback(context.Background(), &inventoryclient.SellInfo{GoodsInfo: items, OrderSn: req.OrderSn})
	if err != nil {
		HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
