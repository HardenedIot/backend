package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"hardenediot-client-service/validator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListHandler(ctx *gin.Context) {
	var results []interface{}
	var err error

	switch ctx.Request.URL.Path {
	case "/users":
		var users []models.User
		err = db.DB.Where("private = ?", false).Find(&users).Error
		results = make([]interface{}, len(users))
		for i, user := range users {
			results[i] = user
		}
	case "/teams":
		var teams []models.Team
		err = db.DB.Where("private = ?", false).Find(&teams).Error
		results = make([]interface{}, len(teams))
		for i, team := range teams {
			results[i] = team
		}
	case "/projects":
		var projects []models.Project
		err = db.DB.Where("private = ?", false).Find(&projects).Error
		results = make([]interface{}, len(projects))
		for i, project := range projects {
			results[i] = project
		}
	default:
		log.Println("Invalid resource path:", ctx.Request.URL.Path)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, results)
}

func CreateHandler(ctx *gin.Context) {
	var model interface{}

	switch ctx.Request.URL.Path {
	case "/users":
		model = &models.User{}
	case "/teams":
		model = &models.Team{}
	case "/projects":
		model = &models.Project{}
	default:
		log.Println("Invalid resource path:", ctx.Request.URL.Path)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := ctx.ShouldBindJSON(model); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := validator.Validate.Struct(model); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := db.DB.Create(model).Error; err != nil {
		log.Println("Error creating instance:", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, model)
}
