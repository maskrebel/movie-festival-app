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

func setupTestRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.POST("/auth/login", controllers.Login(db))
	router.GET("/auth/status", middlewares.AuthMiddleware(db), controllers.Status(db))
	return router
}

func TestRegister(t *testing.T) {
	db := config.SetupDatabase()
	router := setupTestRouter(db)

	// Test case: Successful registration
	t.Run("Successful Login", func(t *testing.T) {
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

		reqGet, err := http.NewRequest("GET", "/auth/status", nil)
		if err != nil {
			t.Errorf("Error in creating GET request: %v", err)
		}
		reqGet.Header.Add("Authorization", "Bearer "+res["token"])
		wGet := httptest.NewRecorder()
		router.ServeHTTP(wGet, reqGet)

		if wGet.Code != http.StatusOK {
			t.Errorf("Response code is %v, wanted %v", w.Code, http.StatusOK)
		}
	})
}
