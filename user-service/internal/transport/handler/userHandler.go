package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-service/internal/core/interface/service"
)

func GetAllUsers(service service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := service.GetAllUsers(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error handling"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, users)
	}
}

func GetUserById(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong id"})
			return
		}
		user, err := userService.GetUserById(c, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "error handling"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, user)
	}
}
