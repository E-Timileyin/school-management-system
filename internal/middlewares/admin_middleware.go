package middlewares

import (
	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		userModel := user.(*models.User)
		if userModel.Role != "admin" {
			c.JSON(403, gin.H{"error": "forbidden - admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
