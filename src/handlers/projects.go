package handlers

import (
	"encoding/json"
	"hardenediot/db"
	"hardenediot/models"
	"hardenediot/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListProjects(ctx *gin.Context) {
	var projects []models.Project
	if err := db.DB.Preload("Team").Find(&projects).Error; err != nil {
		log.Printf("Error fetching projects: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Fetched %d projects", len(projects))
	ctx.JSON(http.StatusOK, projects)
}

func CreateProject(ctx *gin.Context) {
	var project models.Project
	if err := ctx.ShouldBindJSON(&project); err != nil || validate.Struct(project) != nil {
		log.Printf("Invalid project data: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := db.DB.Create(&project).Error; err != nil {
		log.Printf("Error creating project: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp, err := http.Get("https://raw.githubusercontent.com/HardenedIot/data/main/data.json")
	if err != nil {
		log.Printf("Error fetching tasks: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	defer resp.Body.Close()

	var tasks []models.Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for _, jsonTask := range tasks {
		if isTechnologyInProject(jsonTask.Technology, project.Technologies) {
			count, err := storage.DB.Collection(project.ProjectID).CountDocuments(ctx, bson.M{"task_id": jsonTask.TaskID})
			if err != nil {
				log.Printf("Error counting documents: %v", err)
				ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}

			if count > 0 {
				ctx.JSON(http.StatusConflict, http.StatusText(http.StatusConflict))
				return
			}

			_, err = storage.DB.Collection(project.ProjectID).InsertOne(ctx, jsonTask)
			if err != nil {
				log.Printf("Error inserting task: %v", err)
				ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		}
	}

	log.Printf("Project created: %v", project)
	ctx.JSON(http.StatusCreated, project)
}

func isTechnologyInProject(tech models.Technology, technologies models.StringSlice) bool {
	for _, t := range technologies {
		if t == tech {
			return true
		}
	}
	return false
}

func GetProject(ctx *gin.Context) {
	var project models.Project
	if err := db.DB.Preload("Team").Where("project_id = ?", ctx.Param("project_id")).First(&project).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	log.Printf("Fetched project: %v", project)
	ctx.JSON(http.StatusOK, project)
}

func PatchProject(ctx *gin.Context) {
	var project models.Project
	if err := db.DB.Where("project_id = ?", ctx.Param("project_id")).First(&project).Error; err != nil {
		handleDBError(ctx, err)
		return
	}

	var patchRequest models.PatchProjectRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		log.Printf("Invalid patch request: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if patchRequest.ProjectName != nil {
		project.ProjectName = *patchRequest.ProjectName
	}
	if patchRequest.TeamID != nil {
		project.TeamID = *patchRequest.TeamID
	}
	if patchRequest.Private != nil {
		project.Private = patchRequest.Private
	}
	if patchRequest.Description != nil {
		project.Description = *patchRequest.Description
	}
	if patchRequest.URL != nil {
		project.URL = *patchRequest.URL
	}
	if patchRequest.Technologies != nil {
		project.Technologies = *patchRequest.Technologies
	}

	if err := db.DB.Save(&project).Error; err != nil {
		log.Printf("Error updating project: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Project updated: %v", project)
	ctx.JSON(http.StatusOK, project)
}

func DeleteProject(ctx *gin.Context) {
	var project models.Project
	if err := db.DB.Where("project_id = ?", ctx.Param("project_id")).First(&project).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	if err := db.DB.Delete(&project).Error; err != nil {
		log.Printf("Error deleting project: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Project deleted: %v", project)
	ctx.JSON(http.StatusNoContent, nil)
}
