package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *handler) getTransaction(context *gin.Context) {
	id := context.Param("id")

	transaction, err := handler.transactionService.GetByID(id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, transaction)
}
