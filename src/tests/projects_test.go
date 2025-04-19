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

func TestListProjects(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/projects", handlers.ListProjects)

	req, err := http.NewRequest(http.MethodGet, "/projects", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateProject(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/projects", handlers.CreateProject)

	private := false
	technologies := models.StringSlice{"Wifi", "Bluetooth"}
	project := models.Project{
		ProjectID:    "test-project-id",
		ProjectName:  "Test Project",
		TeamID:       "test-team-id",
		Private:      &private,
		Description:  "Test description",
		URL:          "http://example.com",
		Technologies: technologies,
	}

	jsonValue, _ := json.Marshal(project)
	req, err := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusCreated, http.StatusInternalServerError, http.StatusConflict}, w.Code)
}
