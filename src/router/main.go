package routes

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(logger.SetLogger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "http://localhost:3003"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func Run() {
	var r = gin.Default()
	setupRouter(r)
	getRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
