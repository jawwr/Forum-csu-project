package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/internal/core/service"
	"user-service/internal/transport/handler"
	"user-service/internal/transport/middleware"
)

func InitRoutes(serviceManager service.Manager) *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/auth")
	{
		auth.POST("/register", handler.RegisterUser(serviceManager.AuthService))
		auth.POST("/login", handler.LoginUser(serviceManager.AuthService))
	}

	api := router.Group("/api", middleware.AuthMiddleware(serviceManager.TokenService, serviceManager.UserService))
	{
		api.GET("/hello", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Hello"})
		})
	}
	return router
}
