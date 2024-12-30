package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"net/http"
)

func ViewMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		movieID := c.Param("movie_id")

		var input struct {
			Duration int `json:"duration" binding:"required"`
		}

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var movie models.Movie
		if err := db.First(&movie, movieID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found!"})
			return
		}

		view := models.View{
			UserID:   nil,
			MovieID:  movie.ID,
			Duration: input.Duration,
		}

		if userID.(uint) != 0 {
			var id = userID.(uint)
			view.UserID = &id
		}

		if err := db.Create(&view).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update movie.views
		movie.Views += 1
		db.Save(&movie)

		c.JSON(http.StatusOK, gin.H{"message": "View record successfully!"})
	}
}
