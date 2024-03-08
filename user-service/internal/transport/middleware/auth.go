package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"user-service/internal/core/interface/service"
)

func AuthMiddleware(tokenService service.TokenService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization"})
			return
		}

		token, err := tokenService.GetToken(c, auth[len("Bearer "):])
		if err != nil || token.Revoked || token.Expire.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			return
		}

		user, err := userService.GetUserById(c, token.UserId)
		if err != nil {
			log.Fatal("Error during getting user")
		}

		c.Set("user", user)

		c.Next()
	}
}
