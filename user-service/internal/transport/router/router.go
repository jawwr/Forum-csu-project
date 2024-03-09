package router

import (
	"github.com/gin-gonic/gin"
	"user-service/internal/core/service"
	"user-service/internal/transport/handler"
	"user-service/internal/transport/middleware"
)

func InitRoutes(serviceManager service.Manager) *gin.Engine {
	engine := gin.New()

	router := engine.Group("", configureCORS())

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
	return engine
}

func configureCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
