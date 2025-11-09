package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService        domain.UserService
	transactionService domain.TransactionService
}

func NewHandler(userService domain.UserService, transactionService domain.TransactionService) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization"},
	}))

	h := &handler{
		userService:        userService,
		transactionService: transactionService,
	}

	v1 := router.Group("/v1")
	registerUserRoutes(v1, h)
	registerTransactionRoutes(v1, h)

	router.GET("/__health", func(c *gin.Context) { c.String(200, "ok") })

	return router
}

func registerUserRoutes(v1 *gin.RouterGroup, h *handler) {
	v1.POST("/users", h.postUser)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id", h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)
}

func registerTransactionRoutes(v1 *gin.RouterGroup, h *handler) {
	v1.GET("/transactions", h.getTransactions)
	v1.GET("/transactions/:id", h.getTransaction)
	v1.POST("/transactions", h.postTransaction)
	v1.POST("/statement", h.importTransactions)
}
