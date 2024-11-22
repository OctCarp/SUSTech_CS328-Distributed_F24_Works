package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"octcarp/sustech/cs328/a2/api/models"
	"octcarp/sustech/cs328/a2/api/utils"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, models.Message{
				Message: "No token, permission denied",
			})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		userId, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Message{
				Message: "Invalid token, permission denied",
			})
			c.Abort()
			return
		}

		c.Set("token_user_id", userId)
		c.Next()
	}
}
