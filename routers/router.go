package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/setting"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "nihao",
		})

	})
	return engine
}
