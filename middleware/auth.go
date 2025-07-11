package middleware

import (
	"strings"

	"github.com/OctavianoRyan25/VhiWEB/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Need to provide a valid authorization",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		if token == "" {
			c.JSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Need to provide a valid token",
			})
			c.Abort()
			return
		}

		claims, err := util.ValidateJWT(token)
		if err != nil {
			c.JSON(401, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		userIDFloat, okID := claims["user_id"].(float64)
		userEmail, okEmail := claims["email"].(string)

		if !okID || !okEmail {
			c.JSON(401, gin.H{"error": "Invalid token payload"})
			c.Abort()
			return
		}

		c.Set("user_id", int(userIDFloat))
		c.Set("user_email", userEmail)

		c.Next()
	}
}
