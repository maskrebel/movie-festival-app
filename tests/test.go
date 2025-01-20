package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
	"movie-festival-app/models"
	"movie-festival-app/services"
)

func setupDBTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Vote{},
		&models.View{},
		&models.TokenExpired{},
	)

	return db
}

func SetupRouterTest() *gin.Engine {
	// init db
	db := setupDBTest()

	// initial service & controller
	movieService := services.NewMovieService()
	mc := controllers.NewMovieController(db, *movieService)

	// init middleware
	m := middlewares.NewMiddleware(db)

	router := gin.Default()

	// admin
	router.POST("/auth/login", controllers.Login(db))
	router.GET("/admin/most-viewed", m.Auth(), mc.GetMostViewer())
	router.GET("/admin/genre-most-viewed", m.Auth(), mc.GetGenreMostViewed())

	// movie
	router.GET("/movies/search", controllers.SearchMovies(db))
	router.POST("/movies/:id/view", m.Auth(), controllers.ViewMovie(db))

	// auth
	router.POST("/auth/register", controllers.Register(db))
	router.POST("/auth/login", controllers.Login(db))
	router.GET("/auth/status", m.Auth(), controllers.Status(db))

	return router
}
