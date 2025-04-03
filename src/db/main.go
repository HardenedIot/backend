package db

import (
	"hardenediot-client-service/models"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatalln("DB_HOST is not specified")
	}

	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatalln("POSTGRES_USER is not specified")
	}

	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		log.Fatalln("POSTGRES_DB is not specified")
	}

	passwordPath := os.Getenv("POSTGRES_PASSWORD_FILE")
	if passwordPath == "" {
		log.Fatalln("POSTGRES_PASSWORD_FILE is not specified")
	}
	passwordByte, err := os.ReadFile(passwordPath)
	if err != nil {
		log.Fatalln(err)
	}
	password := strings.TrimSpace(string(passwordByte))

	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		log.Fatalln("TIMEZONE is not specified")
	}

	var dsnBuilder strings.Builder
	dsnBuilder.WriteString("host=")
	dsnBuilder.WriteString(host)
	dsnBuilder.WriteString(" user=")
	dsnBuilder.WriteString(user)
	dsnBuilder.WriteString(" password=")
	dsnBuilder.WriteString(password)
	dsnBuilder.WriteString(" dbname=")
	dsnBuilder.WriteString(db)
	dsnBuilder.WriteString(" port=5432 sslmode=disable TimeZone=")
	dsnBuilder.WriteString(timezone)
	dsn := dsnBuilder.String()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return err
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Team{})
	DB.AutoMigrate(&models.Project{})

	return nil
}
