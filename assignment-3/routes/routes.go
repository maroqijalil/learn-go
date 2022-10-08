package routes

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func InitServer(controllers *controllers.Controllers) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router
}
