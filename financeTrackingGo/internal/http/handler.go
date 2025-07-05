package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService        domain.UserService
	transactionService domain.TransactionService
}

func NewClientHandler(userService domain.UserService) http.Handler {
	h := &handler{
		userService: userService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1 := router.Group("/v1")

	v1.POST("/users", h.postUser)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id", h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)

	return router
}

func NewTransactionHandler(transactionService domain.TransactionService) http.Handler {
	h := &handler{
		transactionService: transactionService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/transactions/:id", h.getTransaction)

	return router
}
