package routes

import (
	"hardenediot-client-service/handlers"

	"github.com/gin-gonic/gin"
)

func addRoutes(path string, r *gin.Engine,
	GetHandler gin.HandlerFunc,
	PatchHandler gin.HandlerFunc, DeleteHandler gin.HandlerFunc) {
	group := r.Group(path)

	group.GET(":key", GetHandler)
	group.PATCH(":key", PatchHandler)
	group.DELETE(":key", DeleteHandler)
}

func getRoutes(r *gin.Engine) {
	r.GET("/health", handlers.Health)

	group := r.Group("/auth")
	{
		group.POST("/register", handlers.Health)
		group.POST("/login", handlers.Health)
	}

	r.GET("/users", handlers.ListHandler)
	r.POST("/users", handlers.CreateHandler)
	addRoutes("/users", r,
		handlers.GetHandler,
		handlers.PatchHandler, handlers.DeleteHandler)

	r.GET("/teams", handlers.ListHandler)
	r.POST("/teams", handlers.CreateHandler)
	addRoutes("/teams", r,
		handlers.GetHandler,
		handlers.PatchHandler, handlers.DeleteHandler)

	r.GET("/projects", handlers.ListHandler)
	r.POST("/projects", handlers.CreateHandler)
	addRoutes("/projects", r,
		handlers.GetHandler,
		handlers.PatchHandler, handlers.DeleteHandler)
}
