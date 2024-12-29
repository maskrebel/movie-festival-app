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
		userID, _ := c.Get("user_id")

		var movie models.Movie
		if err := db.First(&movie, movieID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found!"})
			return
		}

		var existingVote models.Vote
		if err := db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(&existingVote).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You have already voted for this movie"})
			return
		}

		input := models.Vote{
			UserID:  userID.(uint),
			MovieID: movie.ID,
		}

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You have already voted for this movie"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Vote successfully!"})
	}
}

func UnVoteMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID := c.Param("id")
		UserID, _ := c.Get("user_id")

		var movie models.Movie
		if err := db.First(&movie, movieID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found!"})
			return
		}

		var vote models.Vote
		if err := db.Where("user_id = ? and movie_id = ?", UserID.(uint), movie.ID).Delete(&vote).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie with this user_id not found!", "d": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "UnVote successfully!"})
	}
}
