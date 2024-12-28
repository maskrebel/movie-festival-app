package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie-festival-app/config"
	"movie-festival-app/routes"
)

func main() {
	db := config.SetupDatabase()

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := gin.Default()
	routes.RegisterRoutes(router, db)

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
