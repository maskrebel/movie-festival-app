package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovies(t *testing.T) {
	router := SetupRouterTest()

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
