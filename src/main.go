package main

import (
	"hardenediot/db"
	routes "hardenediot/router"
	"hardenediot/security"
	"hardenediot/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	security.ReadSecret()
	db.ConnectDB()
	storage.ConnectDB()
	routes.Run()
}
