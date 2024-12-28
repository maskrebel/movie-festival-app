package tests

import (
	"github.com/gin-gonic/gin"
	"movie-festival-app/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	app := gin.New()
	server := httptest.NewServer(app)
	defer server.Close()

	// admin create movie
	resp, err := http.Post(config.BaseUrl()+"/admin/movies", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status code is %d", resp.StatusCode)
	}
}

func TestMovieGetMostViewed(t *testing.T) {
	app := gin.New()
	server := httptest.NewServer(app)
	defer server.Close()

	resp, err := http.Get(config.BaseUrl() + "/admin/movies/most-viewed")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code is %d", resp.StatusCode)
	}
}
