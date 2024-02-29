package http

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/core/interface/service"
	"user-service/internal/transport/handler"
)

func InitRoutes(service service.AuthService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))
	return router
}
