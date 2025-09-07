package http

import (
	"io"
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

func (handler *handler) importTransactions(context *gin.Context) {

	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get file from request" + err.Error(),
		})
		return
	}
	defer file.Close()

	const maxBytes = 10 << 20
	limited := io.LimitReader(file, maxBytes)

	count, err := handler.transactionService.ImportTransactions(limited)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to import transactions: " + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":           "Transactions imported successfully",
		"rows":              count,
		"original_filename": header.Filename,
	})
}
