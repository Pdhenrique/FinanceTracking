package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/gin-gonic/gin"
)

func (handler *handler) getUser(context *gin.Context) {
	id := context.Param("id")

	user, err := handler.userService.GetByID(id)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *handler) postUsers(context *gin.Context) {
	user := &domain.User{}

	if err := context.BindJSON(&user); err != nil {
		return
	}

	user, err := handler.userService.Create(user)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (handler *handler) putUser(context *gin.Context) {}
