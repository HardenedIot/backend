package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(ctx *gin.Context) {
	var users []models.User
	if err := db.DB.Preload("Teams").Where("private = ?", false).Find(&users).Error; err != nil {
		log.Printf("Error fetching users: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("Fetched %d users", len(users))
	ctx.JSON(http.StatusOK, users)
}

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil || validate.Struct(user) != nil {
		log.Printf("Invalid user data: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	if err := db.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("User  created: %v", user)
	ctx.JSON(http.StatusCreated, user)
}

func GetUser(ctx *gin.Context) {
	var user models.User
	if err := db.DB.Preload("Teams").Where("username = ?", ctx.Param("username")).First(&user).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	log.Printf("Fetched user: %v", user)
	ctx.JSON(http.StatusOK, user)
}

func PatchUser(ctx *gin.Context) {
	var user models.User
	if err := db.DB.Where("username = ?", ctx.Param("username")).First(&user).Error; err != nil {
		handleDBError(ctx, err)
		return
	}

	var patchRequest models.PatchUserRequest
	if err := ctx.ShouldBindJSON(&patchRequest); err != nil {
		log.Printf("Invalid patch request: %v", err)
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if patchRequest.Username != nil {
		user.Username = *patchRequest.Username
	}
	if patchRequest.Name != nil {
		user.Name = *patchRequest.Name
	}
	if patchRequest.Surname != nil {
		user.Surname = *patchRequest.Surname
	}
	if patchRequest.Email != nil {
		user.Email = *patchRequest.Email
	}
	if patchRequest.Private != nil {
		user.Private = *patchRequest.Private
	}
	if patchRequest.TeamIDs != nil {
		var teams []models.Team
		if err := db.DB.Where("id IN ?", *patchRequest.TeamIDs).Find(&teams).Error; err != nil {
			log.Printf("Invalid team(s) in patch request: %v", err)
			ctx.JSON(http.StatusBadRequest, "Invalid team(s)")
			return
		}
		user.Teams = teams
	}

	if err := db.DB.Save(&user).Error; err != nil {
		log.Printf("Error updating user: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("User  updated: %v", user)
	ctx.JSON(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {
	var user models.User
	if err := db.DB.Where("username = ?", ctx.Param("username")).First(&user).Error; err != nil {
		handleDBError(ctx, err)
		return
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		log.Printf("Error deleting user: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	log.Printf("User  deleted: %v", user)
	ctx.JSON(http.StatusNoContent, nil)
}
