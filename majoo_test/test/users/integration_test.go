package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/majoo-test/routes"
	"github.com/majoo-test/util"
	"github.com/stretchr/testify/assert"
)

func LoginAccount(t *testing.T) interface{} {
	// Open connection
	db := util.SetupDB()
	router := routes.SetupRouter(db)

	// Data body with data from create account random
	dataBody := fmt.Sprintf(`{"user_name": "%s", "password": "%s"}`, "admin1", "admin1")

	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run server http
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "You have successfully Login", responseBody["message"])
	assert.NotZero(t, responseBody["data"])

	token := responseBody["data"].(map[string]interface{})["token"]
	assert.NotZero(t, token)

	// return token for use any test
	return token
}

// Test Login Success
func TestIntegrationLoginSuccess(t *testing.T) {
	LoginAccount(t)
}

// Test login failed wrong credential
func TestIntegrationLoginWrongCredential(t *testing.T) {
	// Open connection
	db := util.SetupDB()
	router := routes.SetupRouter(db)

	t.Run("wrong password", func(t *testing.T) {
		// Data body with data from create account random
		dataBody := fmt.Sprintf(`{"user_name": "%s", "password": "%s"}`, "admin1", "wrong")

		// Create payload request
		requestBody := strings.NewReader(dataBody)

		// Create request
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
		// Added header content type
		request.Header.Add("Content-Type", "application/json")

		// Create recorder
		recorder := httptest.NewRecorder()

		// Run server http
		router.ServeHTTP(recorder, request)

		// Get response
		response := recorder.Result()

		// Read response
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		// Decode json
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "error", responseBody["status"])
		assert.Equal(t, "Login failed", responseBody["message"])
		assert.NotZero(t, responseBody["data"])
		assert.Equal(t, "email or password incorrect", responseBody["data"].(map[string]interface{})["errors"])
	})

	t.Run("wrong username", func(t *testing.T) {
		// Data body with data from create account random
		dataBody := fmt.Sprintf(`{"user_name": "%s", "password": "%s"}`, "wrong", "admin1")

		// Create payload request
		requestBody := strings.NewReader(dataBody)

		// Create request
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
		// Added header content type
		request.Header.Add("Content-Type", "application/json")

		// Create recorder
		recorder := httptest.NewRecorder()

		// Run server http
		router.ServeHTTP(recorder, request)

		// Get response
		response := recorder.Result()

		// Read response
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		// Decode json
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "error", responseBody["status"])
		assert.Equal(t, "Login failed", responseBody["message"])
		assert.NotZero(t, responseBody["data"])
		assert.Equal(t, "email or password incorrect", responseBody["data"].(map[string]interface{})["errors"])
	})
}

// Test validation error
func TestIntegrationLoginValidationError(t *testing.T) {

	// Open connection
	db := util.SetupDB()
	router := routes.SetupRouter(db)

	// Data body with data from create account random
	dataBody := fmt.Sprintf(`{"user_name": "%s", "password": "%s"}`, "", "wr")

	// Create payload request
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/auth/login", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run server http
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "error", responseBody["status"])
	assert.Equal(t, "Login failed", responseBody["message"])
	assert.NotZero(t, responseBody["data"])
	assert.NotNil(t, responseBody["data"].(map[string]interface{})["errors"])
}
