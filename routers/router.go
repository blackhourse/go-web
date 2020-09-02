package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/setting"
	v1 "go-web/routers/v1"
)

func InitRouter() *gin.Engine {
	/*engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "nihao",
		})

	})
	return engine*/

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/app/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DelTag)
	}
	return r
}
