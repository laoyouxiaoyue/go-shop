package router

import (
	"github.com/gin-gonic/gin"
	"shop/web/api"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	r := Router.Group("order")
	{
		// cart
		r.GET("cart", api.GetCartList)
		r.POST("cart", api.CreateCartItem)
		r.PUT("cart", api.UpdateCartItem)
		r.DELETE("cart", api.DeleteCartItem)

		// order
		r.POST("", api.CreateOrder)
		r.GET("list", api.GetOrderList)
		r.GET("detail/:id", api.GetOrderDetail)
		r.PUT("status", api.UpdateOrderStatus)
	}
}
