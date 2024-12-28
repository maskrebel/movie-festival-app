package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"net/http"
)

func CreateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "error": err.Error()})
		}

		if err := db.Create(&movie).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"is_success": false, "error": err.Error()})
		} else {
			c.JSON(201, gin.H{"is_success": true, "movie_id": movie.ID, "message": "Create movie successfully!"})
		}
	}
}
