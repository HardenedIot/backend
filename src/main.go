package main

import (
	"hardenediot-client-service/db"
	"hardenediot-client-service/routes"
	"hardenediot-client-service/security"
	"hardenediot-client-service/storage"
	"hardenediot-client-service/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	security.ReadSecret()
	db.ConnectDB()
	storage.ConnectDB()
	validator.ValidatorInit()
	routes.Run()
}
