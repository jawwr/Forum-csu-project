package router

import (
	"github.com/gin-gonic/gin"
	"post-service/internal/core/interface/service"
	"post-service/internal/transport/handler"
	"post-service/internal/transport/middleware"
)

func InitRoutes(postService service.PostService, userService service.UserService) *gin.Engine {
	router := gin.New()

	api := router.Group("/api/post", middleware.AuthMiddleware(userService))
	{
		api.POST("", handler.CreatePost(postService))
		api.GET("/:id", handler.GetPost(postService))
		api.GET("", handler.GetAllPosts(postService))
	}

	return router
}
