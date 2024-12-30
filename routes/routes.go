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
		admin.POST("/", middlewares.AuthMiddleware(db), controllers.CreateMovie(db))
		admin.PUT("/:id", middlewares.AuthMiddleware(db), controllers.UpdateMovie(db))
		admin.GET("/most-viewed", middlewares.AuthMiddleware(db), controllers.GetMostViewed(db))
		admin.GET("/genre-most-viewed", middlewares.AuthMiddleware(db), controllers.GetGenreMostViewed(db))
	}

	// auth routes
	auth := router.Group("/auth")
	{
		auth.GET("/status", middlewares.AuthMiddleware(db), controllers.Status(db))
		auth.POST("/register", controllers.Register(db))
		auth.POST("/login", controllers.Login(db))
		auth.POST("/logout", middlewares.AuthMiddleware(db), controllers.Logout(db))
	}

	// movies route
	movies := router.Group("/movies")
	{
		movies.GET("/search", controllers.SearchMovies(db))
		movies.POST("/:id/vote", middlewares.AuthMiddleware(db), controllers.VoteMovie(db))
		movies.DELETE("/:id/unvote", middlewares.AuthMiddleware(db), controllers.UnVoteMovie(db))
		movies.POST("/:id/view", middlewares.AuthMiddleware(db), controllers.ViewMovie(db))
	}
}
