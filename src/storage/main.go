package storage

import (
	"context"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var ctx = context.TODO()

func ConnectDB() error {
	host := os.Getenv("MONGO_HOST")
	if host == "" {
		log.Fatalln("MONGO_HOST is not specified")
	}

	user := os.Getenv("MONGO_USER")
	if user == "" {
		log.Fatalln("MONGO_USER is not specified")
	}

	passwordPath := os.Getenv("MONGO_PASSWORD_FILE")
	if passwordPath == "" {
		log.Fatalln("MONGO_PASSWORD_FILE is not specified")
	}
	passwordByte, err := os.ReadFile(passwordPath)
	if err != nil {
		log.Fatalln(err)
	}
	password := strings.TrimSpace(string(passwordByte))
	if password == "" {
		log.Fatalln("MONGO_PASSWORD is not specified")
	}

	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		log.Fatalln("MONGO_DB is not specified")
	}

	var dsnBuilder strings.Builder
	dsnBuilder.WriteString("mongodb://")
	dsnBuilder.WriteString(user)
	dsnBuilder.WriteString(":")
	dsnBuilder.WriteString(password)
	dsnBuilder.WriteString("@")
	dsnBuilder.WriteString(host)
	dsnBuilder.WriteString(":27017/")
	dsnBuilder.WriteString(dbName)

	clientOptions := options.Client().ApplyURI(dsnBuilder.String())

	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("Connected to Mongo")

	// You can perform any necessary migrations or initializations here

	return nil
}

func DisconnectDB() {
	if err := Client.Disconnect(ctx); err != nil {
		log.Fatalln(err)
	}
	log.Println("Disconnected from MongoDB.")
}
