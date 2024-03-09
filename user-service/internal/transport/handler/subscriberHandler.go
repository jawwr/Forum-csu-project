package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-service/internal/core/interface/service"
	"user-service/internal/transport/model"
)

func SubscribeOnUser(service service.SubscriberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong id"})
			return
		}
		contextUser, _ := c.Get("user")
		user := contextUser.(*model.UserResponse)
		if err := service.Subscribe(c, id, user.Id); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error handling"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Successfully subscribed"})
	}
}

func GetAllSubscribers(service service.SubscriberService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong id"})
			return
		}
		subscribers, err := service.GetAllSubscribers(c, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error handling"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, subscribers)
	}
}
