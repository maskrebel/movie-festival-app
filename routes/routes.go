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
		admin.POST("/", controllers.CreateMovie(db))
	}
}
