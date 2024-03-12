package router

import (
	"github.com/gin-gonic/gin"
	"post-service/internal/core/interface/service"
	"post-service/internal/transport/handler"
	"post-service/internal/transport/middleware"
)

func InitRoutes(postService service.PostService, userService service.UserService) *gin.Engine {
	router := gin.New()

	router.Use(configureCORS())

	api := router.Group("/api/post", middleware.AuthMiddleware(userService))
	{
		api.POST("", handler.CreatePost(postService))
		api.GET("/:id", handler.GetPost(postService, userService))
		api.GET("", handler.GetAllPosts(postService, userService))
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
