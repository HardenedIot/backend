package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"hardenediot/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/health", handlers.Health)

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status":"healthy"}`, w.Body.String())
}
