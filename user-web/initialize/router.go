package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	jz "github.com/zz541843/go-utils"
	"net/http"
	"shop-api/user-web/forms"
	"shop-api/user-web/middlewares"
	"shop-api/user-web/models"
	router2 "shop-api/user-web/router"
	"time"
)

func Routers() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	engine.POST("/login", middlewares.JWTAuth(), func(c *gin.Context) {
		a, _ := c.Get("claims")
		jz.PrintStruct(a)
		c.JSON(http.StatusBadRequest, gin.H{"data": "token验证成功！"})
		/*var json forms.Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	})
	engine.POST("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		j := middlewares.NewJWT()
		claim := models.CustomClaims{
			ID:       1,
			Username: username,
			Password: password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 10,
				Issuer:    "jz",
			},
		}
		token, _ := j.CreateToken(claim)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	engine.POST("/test1", func(c *gin.Context) {
		var u forms.Login
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"success": err.Error(),
			})
			jz.PrintStruct(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true})
	})
	//配置跨域
	//engine.Use(middlewares.Cors())

	ApiGroup := engine.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)

	return engine
}
