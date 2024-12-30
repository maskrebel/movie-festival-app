package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// admin routes
	admin := router.Group("/admin/movies")
	{
		admin.GET("/most-viewed", controllers.GetMostViewed(db))
		admin.POST("/", controllers.CreateMovie(db))
	}

	// auth routes
	auth := router.Group("/auth")
	{
		auth.GET("/status", middlewares.AuthMiddleware(), controllers.Status(db))
		auth.POST("/register", controllers.Register(db))
		auth.POST("/login", controllers.Login(db))
		auth.POST("/logout", controllers.Logout(db))
	}

	// movies route
	movies := router.Group("/movies")
	{
		movies.POST("/:id/vote", middlewares.AuthMiddleware(), controllers.VoteMovie(db))
		movies.DELETE("/:id/unvote", middlewares.AuthMiddleware(), controllers.UnVoteMovie(db))
		movies.POST("/:id/view", middlewares.AuthMiddleware(), controllers.ViewMovie(db))
	}
}
