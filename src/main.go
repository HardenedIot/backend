package main

import (
	"hardenediot-client-service/db"
	routes "hardenediot-client-service/router"
	"hardenediot-client-service/security"
	"hardenediot-client-service/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	security.ReadSecret()
	db.ConnectDB()
	storage.ConnectDB()
	routes.Run()
}
