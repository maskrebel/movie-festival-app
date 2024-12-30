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

func SearchMovies(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := strings.ToLower(c.Query("q"))
		if len(query) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
			return
		}

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 10
		}

		offset := (page - 1) * perPage

		var movies []models.Movie
		if err := db.Where("lower(title) like ? or "+
			"lower(description) like ? or "+
			"? = ANY (string_to_array(lower(artists), ', ')) or "+
			"? = ANY (string_to_array(lower(genres), ', '))", "%"+query+"%", "%"+query+"%", query, query).
			Order("views desc").Offset(offset).Limit(perPage).Find(&movies).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search movies"})
			return
		}

		total := len(movies)
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
				"votes":       movie.Votes,
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
