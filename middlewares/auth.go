package middlewares

import (
	"github.com/gin-gonic/gin"
	"movie-festival-app/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId uint
		authHeader := c.Request.Header.Get("Authorization")
		if !(c.FullPath() == "/movies/:id/view" && authHeader == "") {
			if !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := utils.ParseJWT(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			userId = claims.UserID
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
