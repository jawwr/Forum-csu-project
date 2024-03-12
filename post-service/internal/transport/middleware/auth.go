package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"post-service/internal/core/interface/service"
)

func AuthMiddleware(service service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization"})
			return
		}

		user, err := service.GetUserByToken(c, auth[len("Bearer "):])

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
