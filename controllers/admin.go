package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/models"
	"net/http"
	"strconv"
	"strings"
)

func CreateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Title       string   `json:"title" binding:"required"`
			Year        int      `json:"year" binding:"required"`
			Description string   `json:"description" binding:"required"`
			Duration    int      `json:"duration" binding:"required"`
			Artists     []string `json:"artist"`
			Genres      []string `json:"genre"`
			WatchURL    string   `json:"watch_url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		movie := models.Movie{
			Title:       input.Title,
			Year:        input.Year,
			Description: input.Description,
			Duration:    input.Duration,
			Artists:     strings.Join(input.Artists, ", "),
			Genres:      strings.Join(input.Genres, ", "),
			WatchURL:    input.WatchURL,
		}

		if err := db.Create(&movie).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"movie_id": movie.ID, "message": "Create movie successfully!"})
	}
}

func UpdateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieID, errID := strconv.Atoi(c.Param("id"))
		if errID != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie id"})
			return
		}

		var movie models.Movie
		if err := db.First(&movie, movieID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie not found"})
			return
		}

		var input struct {
			Title       string   `json:"title" binding:"required"`
			Description string   `json:"description" binding:"required"`
			Duration    int      `json:"duration" binding:"required"`
			Artists     []string `json:"artists" binding:"required"`
			Genres      []string `json:"genres" binding:"required"`
			WatchURL    string   `json:"watch_url" binding:"required,url"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// update movie
		movie.Title = input.Title
		movie.Description = input.Description
		movie.Duration = input.Duration
		movie.Artists = strings.Join(input.Artists, ", ")
		movie.Genres = strings.Join(input.Genres, ", ")
		movie.WatchURL = input.WatchURL

		if err := db.Save(&movie).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
	}
}
