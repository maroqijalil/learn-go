package routes

import (
	"assignment-3/controllers"
	"assignment-3/routes/api"
	"assignment-3/routes/web"
	"github.com/gin-gonic/gin"
)

func InitServer(controllers *controllers.Controllers) *gin.Engine {
	router := gin.Default()

	api.InitApi(router, controllers)
	web.InitWeb(router, controllers)

	return router
}
