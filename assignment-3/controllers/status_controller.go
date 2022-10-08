package controllers

import (
	"assignment-3/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller Controllers) GetStatus(context *gin.Context) {
	status, err := repositories.GetStatus()
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, status)
}
