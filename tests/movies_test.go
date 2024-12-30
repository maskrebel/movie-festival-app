package tests

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/config"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestMovieRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/movies/search", controllers.SearchMovies(db))
	router.POST("/movies/:id/view", middlewares.AuthMiddleware(db), controllers.ViewMovie(db))
	return router
}

func TestMovies(t *testing.T) {
	db := config.SetupDatabase()
	router := setupTestMovieRouter(db)

	t.Run("Success Get Movies", func(t *testing.T) {
		query := "dora"
		reqSearch, _ := http.NewRequest("GET", "/movies/search?q="+query, nil)
		wGetSearch := httptest.NewRecorder()

		router.ServeHTTP(wGetSearch, reqSearch)
		if wGetSearch.Code != http.StatusOK {
			t.Errorf("movie search returned wrong status code: got %v want %v", wGetSearch.Code, http.StatusOK)
		}
	})
}
