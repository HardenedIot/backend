package routes

import (
	"hardenediot-client-service/handlers"

	"github.com/gin-gonic/gin"
)

func getRoutes(r *gin.Engine) {
	r.GET("/health", handlers.Health)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Health)
		authGroup.POST("/login", handlers.Health)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("", handlers.ListUsers)
		usersGroup.POST("", handlers.CreateUser)
		usersGroup.GET(":username", handlers.GetUser)
		usersGroup.PATCH(":username", handlers.PatchUser)
		usersGroup.DELETE(":username", handlers.DeleteUser)
	}

	teamsGroup := r.Group("/teams")
	{
		teamsGroup.GET("", handlers.ListTeams)
		teamsGroup.POST("", handlers.CreateTeam)
		teamsGroup.GET(":team_id", handlers.GetTeam)
		teamsGroup.PATCH(":team_id", handlers.PatchTeam)
		teamsGroup.DELETE(":team_id", handlers.DeleteTeam)
	}

	projectsGroup := r.Group("/projects")
	{
		projectsGroup.GET("", handlers.ListProjects)
		projectsGroup.POST("", handlers.CreateProject)
		projectsGroup.GET(":project_id", handlers.GetProject)
		projectsGroup.PATCH(":project_id", handlers.PatchProject)
		projectsGroup.DELETE(":project_id", handlers.DeleteProject)
	}

	tasksGroup := r.Group("/project/:project_id")
	{
		tasksGroup.POST("/init", handlers.Health)
		tasksGroup.GET("/tasks", handlers.GetTasks)
		tasksGroup.PUT("/tasks", handlers.PutTask)
		tasksGroup.PATCH("/tasks", handlers.PatchTask)
		tasksGroup.DELETE("/tasks", handlers.DeleteTask)
	}
}
