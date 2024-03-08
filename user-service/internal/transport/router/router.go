package router

import (
	"github.com/gin-gonic/gin"
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

	api := router.Group("/api/users", middleware.AuthMiddleware(serviceManager.TokenService, serviceManager.UserService))
	{
		api.GET("", handler.GetAllUsers(serviceManager.UserService))

		api.GET("/:id", handler.GetUserById(serviceManager.UserService))

		api.POST("/:id/subscribe", handler.SubscribeOnUser(serviceManager.SubscriberService))

		api.GET("/:id/subscribers", handler.GetAllSubscribers(serviceManager.SubscriberService))
	}
	return router
}
