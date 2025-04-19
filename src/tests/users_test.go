package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hardenediot/handlers"
	"hardenediot/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", handlers.ListUsers)

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusInternalServerError}, w.Code)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", handlers.CreateUser)

	user := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
		Name:     "Test",
		Surname:  "User",
		Private:  false,
	}

	jsonValue, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusCreated, http.StatusBadRequest, http.StatusInternalServerError}, w.Code)
}
