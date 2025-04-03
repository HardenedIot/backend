package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"hardenediot-client-service/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	param := ctx.Param("key")
	var model interface{}

	switch ctx.Request.URL.Path {
	case "/users/" + param:
		model = &models.User{}
	case "/teams/" + param:
		model = &models.Team{}
	case "/projects/" + param:
		model = &models.Project{}
	default:
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	field := "username"
	if ctx.Request.URL.Path[:7] == "/teams/" {
		field = "team_id"
	} else if ctx.Request.URL.Path[:10] == "/projects/" {
		field = "project_id"
	}

	if ctx.Request.URL.Path[:10] == "/projects/" {
		var project models.Project
		if err := db.DB.Where("project_id = ?", param).First(&project).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		ctx.JSON(http.StatusOK, project)
		return
	}

	instance, err := db.FindInstance(db.DB, model, field, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if instance == nil {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	ctx.JSON(http.StatusOK, instance)
}

func DeleteHandler(ctx *gin.Context) {
	param := ctx.Param("key")
	var model interface{}

	switch ctx.Request.URL.Path {
	case "/users/" + param:
		model = &models.User{}
	case "/teams/" + param:
		model = &models.Team{}
	case "/projects/" + param:
		model = &models.Project{}
	default:
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	field := "username"
	if ctx.Request.URL.Path[:7] == "/teams/" {
		field = "team_id"
	} else if ctx.Request.URL.Path[:10] == "/projects/" {
		field = "project_id"
	}

	instance, err := db.FindInstance(db.DB, model, field, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if instance == nil {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if err := db.DeleteInstance(db.DB, model); err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func PatchHandler(ctx *gin.Context) {
	param := ctx.Param("key")
	var model interface{}

	switch ctx.Request.URL.Path {
	case "/users/" + param:
		model = &models.User{}
	case "/teams/" + param:
		model = &models.Team{}
	case "/projects/" + param:
		model = &models.Project{}
	default:
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	field := "username"
	if ctx.Request.URL.Path[:7] == "/teams/" {
		field = "team_id"
	} else if ctx.Request.URL.Path[:10] == "/projects/" {
		field = "project_id"
	}

	instance, err := db.FindInstance(db.DB, model, field, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if instance == nil {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if err := ctx.ShouldBindJSON(model); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := validator.Validate.Struct(model); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := db.UpdateInstance(db.DB, model); err != nil {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model)
}
