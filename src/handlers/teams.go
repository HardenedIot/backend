package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTeams(ctx *gin.Context) {
	var teams []models.Team
	if err := db.DB.Preload("Users").Where("private = ?", false).Find(&teams).Error; err != nil {
		log.Printf("Error fetching teams: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Fetched %d teams", len(teams))
	ctx.JSON(http.StatusOK, teams)
}

func CreateTeam(ctx *gin.Context) {
	var team models.Team
	if err := ctx.ShouldBindJSON(&team); err != nil || validate.Struct(team) != nil {
		log.Printf("Invalid team data: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if err := db.DB.Create(&team).Error; err != nil {
		log.Printf("Error creating team: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Team created: %v", team)
	ctx.JSON(http.StatusCreated, team)
}

func GetTeam(ctx *gin.Context) {
	var team models.Team
	if err := db.DB.Preload("Users").Where("team_id = ?", ctx.Param("team_id")).First(&team).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	log.Printf("Fetched team: %v", team)
	ctx.JSON(http.StatusOK, team)
}

func PatchTeam(ctx *gin.Context) {
	var team models.Team
	if err := db.DB.Where("team_id = ?", ctx.Param("team_id")).First(&team).Error; err != nil {
		handleDBError(ctx, err)
		return
	}

	var patchRequest models.PatchTeamRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		log.Printf("Invalid patch request: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if patchRequest.TeamName != nil {
		team.TeamName = *patchRequest.TeamName
	}
	if patchRequest.TeamID != nil {
		team.TeamID = *patchRequest.TeamID
	}
	if patchRequest.Description != nil {
		team.Description = *patchRequest.Description
	}
	if patchRequest.Users != nil {
		var users []models.User
		if err := db.DB.Where("username IN ?", *patchRequest.Users).Find(&users).Error; err != nil {
			log.Printf("Invalid user(s) in patch request: %v", err)
			ctx.JSON(http.StatusBadRequest, "Invalid user(s)")
			return
		}
		team.Users = users
	}
	if patchRequest.Private != nil {
		team.Private = *patchRequest.Private
	}

	if err := db.DB.Save(&team).Error; err != nil {
		log.Printf("Error updating team: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Team updated: %v", team)
	ctx.JSON(http.StatusOK, team)
}

func DeleteTeam(ctx *gin.Context) {
	var team models.Team
	if err := db.DB.Where("team_id = ?", ctx.Param("team_id")).First(&team).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	if err := db.DB.Delete(&team).Error; err != nil {
		log.Printf("Error deleting team: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Team deleted: %v", team)
	ctx.JSON(http.StatusNoContent, nil)
}
