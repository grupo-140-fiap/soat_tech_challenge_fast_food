package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthHandler(t *testing.T) {
	handler := NewHealthHandler()

	assert.NotNil(t, handler)
	assert.IsType(t, &HealthHandler{}, handler)
}

func TestHealthHandler_HealthCheck_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewHealthHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/health", nil)

	handler.HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestHealthHandler_HealthCheck_ResponseFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewHealthHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/health", nil)

	handler.HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "status")
	assert.Len(t, response, 1)
}

func TestHealthHandler_HealthCheck_DifferentHTTPMethods(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewHealthHandler()

	testCases := []struct {
		name   string
		method string
	}{
		{"GET", http.MethodGet},
		{"POST", http.MethodPost},
		{"PUT", http.MethodPut},
		{"DELETE", http.MethodDelete},
		{"PATCH", http.MethodPatch},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(tc.method, "/health", nil)

			handler.HealthCheck(c)

			assert.Equal(t, http.StatusOK, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "ok", response["status"])
		})
	}
}

func TestHealthHandler_HealthCheck_MultipleConsecutiveCalls(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewHealthHandler()

	for i := 0; i < 5; i++ {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/health", nil)

		handler.HealthCheck(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])
	}
}

func TestHealthHandler_HealthCheck_ResponseBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := NewHealthHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/health", nil)

	handler.HealthCheck(c)

	expectedJSON := `{"status":"ok"}`
	assert.JSONEq(t, expectedJSON, w.Body.String())
}
