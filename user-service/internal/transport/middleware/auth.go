package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"net/http"
	"strings"
	"user-service/internal/core/model"
	"user-service/internal/core/service"
)

func AuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid")
		return
	}

	splitToken := strings.Split(auth, " ")

	login, err := parseToken(splitToken[1])

	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
	}

	c.Set("user", login)

	c.Next()
}

func parseToken(token string) (string, error) {
	// at - access token
	at, err := jwt.ParseWithClaims(token, &service.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid method")
			}

			return []byte(model.SignInKey), nil
		})

	if err != nil {
		return "", err
	}

	claims, ok := at.Claims.(*service.TokenClaims)

	if !ok {
		return "", err
	}

	return claims.Login, nil
}
