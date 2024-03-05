package router

import (
	"github.com/gin-gonic/gin"
	"post-service/internal/core/interface/service"
	"post-service/internal/transport/handler"
)

func InitRoutes(postService service.PostService) *gin.Engine {
	router := gin.New()

	api := router.Group("/api/post")
	{
		api.POST("", handler.CreatePost(postService))
		api.GET("/:id", handler.GetPost(postService))
		api.GET("", handler.GetAllPosts(postService))
	}
	return router
}
