package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/services"
	"net/http"
	"strconv"
)

type MovieController struct {
	db           *gorm.DB
	movieService services.MovieService
}

func NewMovieController(db *gorm.DB, movieService services.MovieService) *MovieController {
	return &MovieController{
		db:           db,
		movieService: movieService,
	}
}

func (mc *MovieController) GetMostViewer() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
		if page < 1 {
			page = 1
		}
		if perPage < 1 {
			perPage = 10
		}
		params := map[string]int{
			"page":     page,
			"per_page": perPage,
		}

		err, data := mc.movieService.GetMostViewed(mc.db, params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (mc *MovieController) GetGenreMostViewed() gin.HandlerFunc {
	return func(c *gin.Context) {
		mapGenresTotal, err := mc.movieService.GetGenreMostViewed(mc.db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": mapGenresTotal})
	}
}
