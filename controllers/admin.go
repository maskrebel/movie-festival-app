package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"movie-festival-app/models"
	"net/http"
	"strconv"
	"strings"
)

func GetMostViewed(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movies []models.Movie
		var total int64

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 10
		}

		db.Model(&models.Movie{}).Count(&total)

		offset := (page - 1) * perPage
		db.Order("views desc").Offset(offset).Limit(perPage).Find(&movies)

		totalPages := int(math.Ceil(float64(total) / float64(perPage)))
		data := make([]map[string]interface{}, 0)
		for _, movie := range movies {
			obj := map[string]interface{}{
				"id":          movie.ID,
				"title":       movie.Title,
				"year":        movie.Year,
				"description": movie.Description,
				"duration":    movie.Duration,
				"artists":     movie.Artists,
				"genres":      movie.Genres,
				"views":       movie.Views,
				"votes":       movie.Views,
				"watch_url":   movie.WatchURL,
				"created_at":  movie.CreatedAt,
				"updated_at":  movie.UpdatedAt,
			}

			data = append(data, obj)
		}

		res := gin.H{
			"page":        page,
			"per_page":    perPage,
			"total":       total,
			"total_pages": totalPages,
			"data":        data,
		}

		c.JSON(http.StatusOK, res)
	}
}

func GetGenreMostViewed(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movies []models.Movie
		mapGenresTotal := make(map[string]int32)

		db.Find(&movies)

		for _, movie := range movies {
			genres := strings.Split(movie.Genres, ", ")
			for _, v := range genres {
				if _, ok := mapGenresTotal[v]; ok {
					mapGenresTotal[v] += int32(movie.Views)
					continue
				}
				mapGenresTotal[v] = int32(movie.Views)
			}
		}

		c.JSON(http.StatusOK, gin.H{"data": mapGenresTotal})
	}
}

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
