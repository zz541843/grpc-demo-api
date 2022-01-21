package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/api"
	"shop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", middlewares.JWTAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)

		UserRouter.GET("detail", middlewares.JWTAuth(), api.GetUserDetail)
		UserRouter.PATCH("update", middlewares.JWTAuth(), api.UpdateUser)
	}
	//服务注册和发现
}
