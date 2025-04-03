package routes

import (
	"log"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(logger.SetLogger())
}

func Run() {
	var r = gin.Default()
	setupRouter(r)
	getRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
