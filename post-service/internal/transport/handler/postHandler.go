package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"post-service/internal/core/interface/service"
	"post-service/internal/core/model"
	"strconv"
)

type handlerPost struct {
	Id    int        `json:"id"`
	Title string     `json:"title"`
	Text  string     `json:"text"`
	User  handleUser `json:"user"`
}

type handleUser struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

func CreatePost(service service.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post model.Post

		if err := c.BindJSON(&post); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}
		ctxUser, exist := c.Get("user")
		if !exist {
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

		c.JSON(http.StatusOK, handlerPost{
			Id:    post.Id,
			Title: post.Title,
			Text:  post.Text,
			User: handleUser{
				Id:    post.UserId,
				Login: "",
			},
		})
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
				Id:    post.Id,
				Title: post.Title,
				Text:  post.Text,
				User: handleUser{
					Id:    post.UserId,
					Login: "",
				},
			})
		}
		c.JSON(http.StatusOK, result)
	}
}
