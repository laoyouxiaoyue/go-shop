package initialize

import (
	"github.com/gin-gonic/gin"
	"shop/web/middlewares"
	router2 "shop/web/router"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.JWTAuth())
	ApiGroup := router.Group("/v1")
	router2.InitUserRouter(ApiGroup)
	return router
}
