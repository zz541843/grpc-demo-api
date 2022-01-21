package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	router2 "shop-api/user-web/router"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	//engine.Use(middlewares.Cors())

	ApiGroup := engine.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)

	return engine
}
