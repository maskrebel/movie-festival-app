package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"movie-festival-app/config"
	"movie-festival-app/models"
	"movie-festival-app/routes"
	"os"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Movie{})
}

func runServer(db *gorm.DB) {
	router := gin.Default()
	routes.RegisterRoutes(router, db)

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
