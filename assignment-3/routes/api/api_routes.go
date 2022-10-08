package api

import (
	"assignment-3/controllers"
	"github.com/gin-gonic/gin"
)

func InitApi(router *gin.Engine, controllers *controllers.Controllers) {
	api := router.Group("/api")
	{
		api.Use(func() gin.HandlerFunc {
			return func(context *gin.Context) {
				context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				context.Next()
			}
		}())
		api.GET("/status", controllers.GetStatus)
	}
}
