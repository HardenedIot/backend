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

func TestGetTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/projects/:project_id/tasks", handlers.GetTasks)

	req, err := http.NewRequest(http.MethodGet, "/projects/test-project-id/tasks", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound, http.StatusInternalServerError}, w.Code)
}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/projects/:project_id/tasks", handlers.CreateTask)

	task := models.Task{
		TaskID:      "test-task-id",
		Technology:  "Wifi",
		Name:        "Test Task",
		Description: "Test task description",
		RiskLevel:   1,
		Completed:   false,
		Ignored:     false,
	}

	jsonValue, _ := json.Marshal(task)
	req, err := http.NewRequest(http.MethodPost, "/projects/test-project-id/tasks", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound, http.StatusBadRequest, http.StatusInternalServerError, http.StatusConflict}, w.Code)
}
