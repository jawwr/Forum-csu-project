package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/mapper"
	"post-service/internal/core/model"
	"strconv"
)

func CreatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Post

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		ctxUser, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "wrong user"})
			return
		}

		user := ctxUser.(model.User)
		post.UserId = user.Id

		id, err := service.CreatePost(c.Request.Context(), post)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": id})
	}
}

func GetPost(postService service.PostService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		post, err := postService.GetPost(c.Request.Context(), numberId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		user, err := userService.GetUserById(c.Request.Context(), post.UserId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, *mapper.ToHandlerPost(&post, &user))
	}
}

func GetAllPosts(postService service.PostService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := postService.GetAllPosts(c.Request.Context())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		var result []model.HandlePost
		for _, post := range posts {
			user, err := userService.GetUserById(c.Request.Context(), post.UserId)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
				return
			}
			result = append(result, *mapper.ToHandlerPost(&post, &user))
		}
		c.JSON(http.StatusOK, result)
	}
}
