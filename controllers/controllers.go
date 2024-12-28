package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"movie-festival-app/models"
	"net/http"
	"strconv"
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

		res := gin.H{
			"page":        page,
			"per_page":    perPage,
			"total":       total,
			"total_pages": totalPages,
			"data":        movies,
		}

		c.JSON(http.StatusOK, res)
	}
}

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
