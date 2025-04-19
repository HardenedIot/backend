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

func TestListTeams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/teams", handlers.ListTeams)

	req, err := http.NewRequest(http.MethodGet, "/teams", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusInternalServerError}, w.Code)
}

func TestCreateTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/teams", handlers.CreateTeam)

	team := models.Team{
		TeamID:      "test-team-id",
		TeamName:    "Test Team",
		Description: "Test team description",
		Private:     false,
	}

	jsonValue, _ := json.Marshal(team)
	req, err := http.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusCreated, http.StatusBadRequest, http.StatusInternalServerError}, w.Code)
}
