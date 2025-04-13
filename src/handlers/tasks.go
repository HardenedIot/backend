package handlers

import (
	"hardenediot-client-service/models"
	"hardenediot-client-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	collection := storage.DB.Collection(projectID)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
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

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	collection := storage.DB.Collection(projectID)

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func PatchTask(ctx *gin.Context) {
	projectID := ctx.Param("project_id")

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	collection := storage.DB.Collection(projectID)
	filter := bson.M{"task_id": task.TaskID}
	update := bson.M{"$set": task}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func DeleteTask(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	collection := storage.DB.Collection(projectID)
	filter := bson.M{"task_id": task.TaskID}

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
