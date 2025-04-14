package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func handleDBError(ctx *gin.Context, err error) {
	log.Printf("Database error: %v", err)
	if err.Error() == "record not found" {
		ctx.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
