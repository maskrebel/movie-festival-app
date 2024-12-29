package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/controllers"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// admin routes
	admin := router.Group("/admin/movies")
	{
		admin.GET("/most-viewed", controllers.GetMostViewed(db))
		admin.POST("/", controllers.CreateMovie(db))
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register(db))
		auth.POST("/login", controllers.Login(db))
		auth.POST("/logout", controllers.Logout(db))
	}
}
