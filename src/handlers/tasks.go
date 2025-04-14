package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"hardenediot-client-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func projectExists(projectID string) bool {
	var count int64
	if err := db.DB.Model(&models.Project{}).Where("project_id = ?", projectID).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func GetTasks(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if !projectExists(projectID) {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	collection := storage.DB.Collection(projectID)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func PutTask(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if !projectExists(projectID) {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil || validate.Struct(task) != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	count, err := storage.DB.Collection(projectID).CountDocuments(ctx, bson.M{"task_id": task.TaskID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if count > 0 {
		ctx.JSON(http.StatusConflict, http.StatusText(http.StatusConflict))
		return
	}

	_, err = storage.DB.Collection(projectID).InsertOne(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func PatchTask(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if !projectExists(projectID) {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	var patchRequest models.PatchTaskRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var task models.Task
	if err := storage.DB.Collection(projectID).FindOne(ctx, bson.M{"task_id": patchRequest.TaskID}).Decode(&task); err != nil {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if patchRequest.TaskID != nil {
		task.TaskID = *patchRequest.TaskID
	}
	if patchRequest.Technology != nil {
		task.Technology = *patchRequest.Technology
	}
	if patchRequest.Name != nil {
		task.Name = *patchRequest.Name
	}
	if patchRequest.Description != nil {
		task.Description = *patchRequest.Description
	}
	if patchRequest.RiskLevel != nil {
		task.RiskLevel = *patchRequest.RiskLevel
	}
	if patchRequest.Completed != nil {
		task.Completed = *patchRequest.Completed
	}

	_, err := storage.DB.Collection(projectID).UpdateOne(ctx, bson.M{"task_id": task.TaskID}, bson.M{"$set": task})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func DeleteTask(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if !projectExists(projectID) {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	var request struct {
		TaskID string `json:"task_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	taskUUID, err := uuid.Parse(request.TaskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	collection := storage.DB.Collection(projectID)
	filter := bson.M{"task_id": taskUUID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
