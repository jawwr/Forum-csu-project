package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"user-service/internal/core/interface/service"
	"user-service/internal/transport/model"
)

func RegisterUser(service service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.UserCredentials

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		token, err := service.Register(c.Request.Context(), user)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func LoginUser(service service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.UserCredentials

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		token, err := service.Login(c, user)

		if err != nil {
			log.Printf("Error during login: %s", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"message": "Wrong credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
