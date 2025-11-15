package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/gin-gonic/gin"
)

func (handler *handler) getAccount(context *gin.Context) {
	id := context.Param("id")

	account, err := handler.accountService.GetByID(id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (handler *handler) postAccount(context *gin.Context) {
	account := &domain.Account{}

	if err := context.BindJSON(&account); err != nil {
		return
	}

	account, err := handler.accountService.Post(account)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, account)

}

func (handler *handler) putAccount(context *gin.Context) {
	id := context.Param("id")
	account := &domain.Account{}

	if err := context.ShouldBindJSON(&account); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "JSON invalido"})

		return
	}

	account.ID = id

	err := handler.accountService.Put(account)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (handler *handler) deleteAccount(context *gin.Context) {
	id := context.Param("id")

	err := handler.accountService.Delete(id)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.AbortWithStatus(http.StatusNoContent)
}
