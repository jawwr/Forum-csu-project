package router

import (
	"event-service/internal/core/interface/service"
	"event-service/internal/transport/handler"
	"event-service/internal/transport/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(eventService service.EventService, userService service.UserService) *gin.Engine {
	router := gin.New()

	router.Use(configureCORS())

	api := router.Group("/api/event", middleware.AuthMiddleware(userService))
	{
		api.GET("", handler.GetEvents(eventService))
	}

	return router
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
