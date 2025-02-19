package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"movie-festival-app/utils"
	"net/http"
	"strings"
)

type Middleware struct {
	db *gorm.DB
}

func NewMiddleware(db *gorm.DB) *Middleware {
	return &Middleware{
		db: db,
	}
}

func (am *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId uint
		var username string
		var tokenString string
		authHeader := c.Request.Header.Get("Authorization")
		if !(c.FullPath() == "/movies/:id/view" && authHeader == "") {
			if !strings.HasPrefix(authHeader, "Bearer ") {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			tokenString = strings.TrimPrefix(authHeader, "Bearer ")

			// validate logout
			var tokenExpired models.TokenExpired
			if err := am.db.Where(&models.TokenExpired{Token: tokenString}).First(&tokenExpired).Error; err == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Your token is expired"})
				c.Abort()
				return
			}

			claims, err := utils.ParseJWT(tokenString)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			userId = claims.UserID

			// get from database
			var user models.User
			am.db.First(&user, userId)

			username = user.Username
			if strings.Split(c.FullPath(), "/")[1] == "admin" && !user.IsAdmin {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission to view this resource"})
				c.Abort()
				return
			}
		}
		c.Set("user_id", userId)
		c.Set("username", username)
		c.Set("token", tokenString)
		c.Next()
	}
}
