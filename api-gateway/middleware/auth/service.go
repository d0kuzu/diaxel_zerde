package auth

import (
	"api-gateway/grpc/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ServiceMiddleware(db *db.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid Authorization format",
			})
			return
		}

		tokenStr := parts[1]

		assistant, err := db.GetAssistant(tokenStr) //TODO: заменить на поиск по токену
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token not found",
			})
			return
		}

		c.Set("user_id", assistant.UserId)
		c.Set("assistant_id", assistant.Id)

		c.Next()
	}
}
