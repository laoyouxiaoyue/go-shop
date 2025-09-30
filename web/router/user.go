package router

import (
	"github.com/gin-gonic/gin"
	"shop/web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("old_user")
	{
		UserRouter.GET("list", api.GetUserList)
		UserRouter.POST("login/password", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
