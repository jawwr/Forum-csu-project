package handler

import (
	"event-service/internal/core/interface/service"
	"event-service/internal/core/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEvents(eventService service.EventService) gin.HandlerFunc {
	return func(c *gin.Context) {
		savedUser, exists := c.Get("user")

		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				`message`: `error saving user`,
			})

			return
		}

		user := savedUser.(model.User)

		events, err := eventService.GetUserEvents(c, user.Id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				`message`: `error saving user`,
			})

			return
		}

		c.AbortWithStatusJSON(http.StatusOK, events)
	}
}
