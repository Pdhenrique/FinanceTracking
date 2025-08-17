package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
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

func (handler *handler) postTransaction(context *gin.Context) {
	transaction := &domain.Transaction{}

	if err := context.BindJSON(&transaction); err != nil {
		return
	}

	transaction, err := handler.transactionService.Post(transaction)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, transaction)

}
