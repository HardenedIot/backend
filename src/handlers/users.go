package handlers

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/models"
	"hardenediot-client-service/security"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func RegisterUser(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	hashedPassword := security.GeneratePasswordHash(input.Password)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Name:     input.Name,
		Surname:  input.Surname,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	hashedInputPassword := security.GeneratePasswordHash(input.Password)
	if hashedInputPassword != user.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Println("SECRET not configured")
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Failed to generate token")
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
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
