package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
