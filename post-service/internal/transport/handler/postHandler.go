package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
	"strconv"
)

type handlerPost struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	UserId int    `json:"user_id"`
}

func CreatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post handlerPost

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		id, err := service.CreatePost(c.Request.Context(), model.Post(post))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": id})
	}
}

func GetPost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id поста"})

			return
		}

		post, err := service.GetPost(c.Request.Context(), numberId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return

		}

		c.JSON(http.StatusOK, handlerPost(post))
	}
}

func GetAllPosts(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := service.GetAllPosts(c.Request.Context())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		var result []handlerPost
		for _, post := range posts {
			result = append(result, handlerPost{
				Id:     post.Id,
				Text:   post.Text,
				UserId: post.UserId,
			})
		}
		c.JSON(http.StatusOK, result)
	}
}
