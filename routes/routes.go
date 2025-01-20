package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, mc controllers.MovieController, m middlewares.Middleware) {
	// admin routes
	admin := router.Group("/admin/movies")
	{
		admin.POST("/", m.Auth(), controllers.CreateMovie(db))
		admin.PUT("/:id", m.Auth(), controllers.UpdateMovie(db))
		admin.GET("/most-viewed", m.Auth(), mc.GetMostViewer())
		admin.GET("/genre-most-viewed", m.Auth(), mc.GetGenreMostViewed())

		admin.GET("/v1/genre-most-viewed", m.Auth(), mc.GetGenreMostViewed())
	}

	// auth routes
	auth := router.Group("/auth")
	{
		auth.GET("/status", m.Auth(), controllers.Status(db))
		auth.POST("/register", controllers.Register(db))
		auth.POST("/login", controllers.Login(db))
		auth.POST("/logout", m.Auth(), controllers.Logout(db))
	}

	// movies route
	movies := router.Group("/movies")
	{
		movies.GET("/search", controllers.SearchMovies(db))
		movies.POST("/:id/vote", m.Auth(), controllers.VoteMovie(db))
		movies.DELETE("/:id/unvote", m.Auth(), controllers.UnVoteMovie(db))
		movies.POST("/:id/view", m.Auth(), controllers.ViewMovie(db))
	}
}
