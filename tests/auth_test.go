package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	router := SetupRouterTest()

	t.Run("Successful Register", func(t *testing.T) {
		requestBody := map[string]string{
			"username": "test_acc",
			"email":    "acc@test.com",
			"password": "securepassword",
		}

		body, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
		req.Header.Add("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})

	// Test case: Successful registration
	t.Run("Successful Login", func(t *testing.T) {
		requestBody := map[string]string{
			"email":    "acc@test.com",
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
