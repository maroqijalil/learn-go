package controllers

import (
	"assignment-2/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller *Controllers) CreateOrder(context *gin.Context) {
	var order models.Order
	if err := context.ShouldBindJSON(&order); err != nil {
		log.Println("bad request")
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	database := controller.Database
	if err := database.
		Select("Items").
		Create(&order).Error; err != nil {
		log.Println("failed to insert order into database")
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//if err := database.
	//	Model(&models.Order{OrderId: order.OrderId}).
	//	Association("Items").
	//	Append(order.Items); err != nil {
	//	log.Println("failed to insert items into database")
	//	context.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}

	log.Println("succeeded to insert into database")
	context.JSON(http.StatusCreated, gin.H{
		"order": order,
	})
}

func (controller *Controllers) GetOrders(context *gin.Context) {
	database := controller.Database
	var orders []models.Order
	if err := database.
		Model(&models.Order{}).
		Preload("Items").
		Find(&orders).Error; err != nil {
		log.Println("failed to get data from database")
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("succeeded to get data from database")
	context.JSON(http.StatusCreated, gin.H{
		"orders": orders,
	})
}

func (controller *Controllers) UpdateOrder(context *gin.Context) {
	var orderId uint
	if status := handleOrderIdParam(context, &orderId); !status {
		return
	}

	var newOrder models.Order
	if err := context.ShouldBindJSON(&newOrder); err != nil {
		log.Println("bad request")
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	database := controller.Database
	if err := database.
		Model(&models.Order{OrderId: orderId}).
		Updates(newOrder).Error; err != nil {
		log.Println("failed to update order of the record")
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, item := range newOrder.Items {
		if err := database.
			Model(&models.Item{ItemId: item.ItemId}).
			Where("order_id = ?", orderId).
			Omit("OrderId").
			Updates(item).Error; err != nil {
			log.Println("failed to update items of into record")
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	log.Println("succeeded to update the record")
	context.JSON(http.StatusCreated, gin.H{
		"message": "The record is successfully updated",
	})
}

func (controller *Controllers) DeleteOrder(context *gin.Context) {
	var orderId uint
	if status := handleOrderIdParam(context, &orderId); !status {
		return
	}

	database := controller.Database
	if err := database.
		Select("Items").
		Delete(&models.Order{OrderId: orderId}).Error; err != nil {
		log.Println("failed to delete the record")
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Println("succeeded to delete the record")
	context.JSON(http.StatusCreated, gin.H{
		"message": "The record is successfully deleted",
	})
}

func handleOrderIdParam(context *gin.Context, orderId *uint) bool {
	if id, err := strconv.Atoi(context.Param("orderId")); err == nil {
		*orderId = uint(id)
		return true
	} else {
		log.Println("bad parameter")
		context.AbortWithError(http.StatusBadRequest, err)
		return false
	}
}
