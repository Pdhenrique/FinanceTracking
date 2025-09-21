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

func NewHandler(userService domain.UserService, transactionService domain.TransactionService) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	h := &handler{
		userService:        userService,
		transactionService: transactionService,
	}

	v1 := r.Group("/v1")
	registerUserRoutes(v1, h)
	registerTransactionRoutes(v1, h)

	r.GET("/__health", func(c *gin.Context) { c.String(200, "ok") })

	return r
}

func registerUserRoutes(v1 *gin.RouterGroup, h *handler) {
	v1.POST("/users", h.postUser)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id", h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)
}

func registerTransactionRoutes(v1 *gin.RouterGroup, h *handler) {
	v1.GET("/transactions/:id", h.getTransaction)
	v1.POST("/transactions", h.postTransaction)
	v1.POST("/statement", h.importTransactions)
}
