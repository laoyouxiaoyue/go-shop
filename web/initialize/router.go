package initialize

import (
	"github.com/gin-gonic/gin"
	"shop/web/middlewares"
	router2 "shop/web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JWTAuth())
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	router2.InitGoodsRouter(ApiGroup)
	return router
}
