package routes

import (
	"hardenediot/handlers"

	"hardenediot/middleware"

	"github.com/gin-gonic/gin"
)

func getRoutes(r *gin.Engine) {
	r.GET("/health", handlers.Health)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.RegisterUser)
		authGroup.POST("/login", handlers.LoginUser)
	}

	usersGroup := r.Group("/users")
	usersGroup.Use(middleware.JWTAuthMiddleware())
	{
		usersGroup.GET("", handlers.ListUsers)
		usersGroup.POST("", handlers.CreateUser)
		usersGroup.GET(":username", handlers.GetUser)
		usersGroup.PATCH(":username", handlers.PatchUser)
		usersGroup.DELETE(":username", handlers.DeleteUser)
	}

	teamsGroup := r.Group("/teams")
	teamsGroup.Use(middleware.JWTAuthMiddleware())
	{
		teamsGroup.GET("", handlers.ListTeams)
		teamsGroup.POST("", handlers.CreateTeam)
		teamsGroup.GET(":team_id", handlers.GetTeam)
		teamsGroup.PATCH(":team_id", handlers.PatchTeam)
		teamsGroup.DELETE(":team_id", handlers.DeleteTeam)
	}

	projectsGroup := r.Group("/projects")
	projectsGroup.Use(middleware.JWTAuthMiddleware())
	{
		projectsGroup.GET("", handlers.ListProjects)
		projectsGroup.POST("", handlers.CreateProject)
		projectsGroup.GET(":project_id", handlers.GetProject)
		projectsGroup.PATCH(":project_id", handlers.PatchProject)
		projectsGroup.DELETE(":project_id", handlers.DeleteProject)
	}

	tasksGroup := r.Group("/project")
	tasksGroup.Use(middleware.JWTAuthMiddleware())
	{
		tasksGroup.GET(":project_id", handlers.GetTasks)
		tasksGroup.POST(":project_id", handlers.CreateTask)
		tasksGroup.PATCH(":project_id", handlers.PatchTask)
		tasksGroup.DELETE(":project_id", handlers.DeleteTask)
	}

	r.GET("/technologies", handlers.ListTechnologies)
}
