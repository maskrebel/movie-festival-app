package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"movie-festival-app/config"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
	"movie-festival-app/models"
	"movie-festival-app/routes"
	"movie-festival-app/services"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Vote{},
		&models.View{},
		&models.TokenExpired{},
	)
}

func runServer(db *gorm.DB) {
	router := gin.Default()

	// initial service & controller
	movieService := services.NewMovieService()
	movieController := controllers.NewMovieController(db, *movieService)

	// initial middleware
	middleware := middlewares.NewMiddleware(db)

	routes.RegisterRoutes(router, db, *movieController, *middleware)

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	if len(os.Args) > 1 {
		db := config.SetupDatabase()

		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		switch os.Args[1] {
		case "run":
			runServer(db)
		case "migrate":
			migrate(db)
		default:
			log.Println("command not recognized")
		}
	}
}
