package web

import (
	"assignment-3/controllers"
	"github.com/gin-gonic/gin"
)

func InitWeb(router *gin.Engine, controllers *controllers.Controllers) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", controllers.Index)
}
