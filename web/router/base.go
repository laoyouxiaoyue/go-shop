package router

import (
	"github.com/gin-gonic/gin"
	"shop/web/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.GET("sendsms")
	}
}
