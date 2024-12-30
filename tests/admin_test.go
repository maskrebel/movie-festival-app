package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"movie-festival-app/config"
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupAdminRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.POST("/auth/login", controllers.Login(db))
	router.GET("/admin/most-viewed", middlewares.AuthMiddleware(db), controllers.GetMostViewed(db))
	router.GET("/admin/genre-most-viewed", middlewares.AuthMiddleware(db), controllers.GetGenreMostViewed(db))
	return router
}

func TestAdmin(t *testing.T) {
	db := config.SetupDatabase()
	router := SetupAdminRouter(db)

	t.Run("Test Admin", func(t *testing.T) {
		requestBody := map[string]string{
			"email":    "user1@example.com",
			"password": "securepassword",
		}
		body, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Response code is %v, wanted %v", w.Code, http.StatusOK)
		}

		var res map[string]string
		json.Unmarshal(w.Body.Bytes(), &res)

		if res["token"] == "" {
			t.Errorf("Token is empty, wanted %v", res["token"])
			return
		}

		reqGetMost, _ := http.NewRequest("GET", "/admin/genre-most-viewed", nil)
		reqGetMost.Header.Add("Authorization", "Bearer "+res["token"])

		wGetMost := httptest.NewRecorder()
		router.ServeHTTP(wGetMost, reqGetMost)

		if wGetMost.Code != http.StatusOK {
			t.Errorf("Response code is %v, wanted %v", wGetMost.Code, http.StatusOK)
		}
	})
}
