package services

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"movie-festival-app/models"
	"strings"
)

type MovieService struct{}

func NewMovieService() *MovieService {
	return new(MovieService)
}

func (ms *MovieService) GetMostViewed(db *gorm.DB, params map[string]int) (error, map[string]interface{}) {
	var movies []models.Movie
	var total int64

	page := params["page"]
	perPage := params["per_page"]

	db.Model(&models.Movie{}).Count(&total)
	if total == 0 {
		return fmt.Errorf("movies is empty"), nil
	}

	offset := (page - 1) * perPage
	db.Order("views desc").Offset(offset).Limit(perPage).Find(&movies)

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	res := map[string]interface{}{
		"page":        page,
		"per_page":    perPage,
		"total":       total,
		"total_pages": totalPages,
		"data":        movies,
	}

	return nil, res
}

func (ms *MovieService) GetGenreMostViewed(db *gorm.DB) (map[string]int, error) {
	mapGenresTotal := make(map[string]int)
	var movies []models.Movie

	db.Find(&movies)
	if movies == nil {
		return nil, fmt.Errorf("movies is empty")
	}

	for _, movie := range movies {
		genres := strings.Split(movie.Genres, ", ")
		for _, v := range genres {
			if _, ok := mapGenresTotal[v]; ok {
				mapGenresTotal[v] += movie.Views
				continue
			}
			mapGenresTotal[v] = movie.Views
		}
	}

	return mapGenresTotal, nil
}
