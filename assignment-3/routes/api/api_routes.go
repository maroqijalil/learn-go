package api

import (
	"assignment-3/controllers"
	"github.com/gin-gonic/gin"
)

func InitApi(router *gin.Engine, controllers *controllers.Controllers) {
	api := router.Group("/api")
	{
		api.GET("/status", controllers.GetStatus)
	}
}
