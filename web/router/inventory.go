package router

import (
	"github.com/gin-gonic/gin"
	"shop/web/api"
)

func InitInventoryRouter(Router *gin.RouterGroup) {
	inv := Router.Group("inventory")
	{
		inv.POST("set", api.SetInventory)
		inv.GET("detail/:goodsId", api.GetInventoryDetail)
		inv.POST("sell", api.SellInventory)
		inv.POST("reback", api.RebackInventory)
	}
}
