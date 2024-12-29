package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"net/http"
)

func VoteMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID := c.Param("id")
		UserID, _ := c.Get("user_id")

		var movie models.Movie
		if err := db.First(&movie, movieID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found!"})
		}

		input := models.Vote{
			UserID:  UserID.(uint),
			MovieID: movie.ID,
		}

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You have already voted for this movie"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Vote successfully!"})
	}
}
