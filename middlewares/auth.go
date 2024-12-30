package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"movie-festival-app/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId uint
		var username string
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

			// get from database
			var user models.User
			db.First(&user, userId)

			username = user.Username
			if strings.Split(c.FullPath(), "/")[1] == "admin" && !user.IsAdmin {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission to view this resource"})
				c.Abort()
				return
			}
		}
		c.Set("user_id", userId)
		c.Set("username", username)
		c.Next()
	}
}
