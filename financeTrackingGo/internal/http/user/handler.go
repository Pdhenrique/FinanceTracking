package http

import (
	"net/http"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService domain.UserService
}

func NewHandler(userService domain.UserService) http.Handler {
	h := &handler{
		userService: userService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1 := router.Group("/v1")

	v1.POST("/users")
	v1.GET("/users/:id")
	v1.PUT("/users/:id")
	v1.DELETE("/users/:id")

	return router
}
