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

	v1.POST("/users", h.postUser)
	v1.GET("/users/:id", h.getUser)
	v1.PUT("/users/:id" h.putUser)
	v1.DELETE("/users/:id", h.deleteUser)

	return router
}
