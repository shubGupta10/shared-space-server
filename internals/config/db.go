package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shubGupta10/shared-space-server/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// Implementation for connecting to the database
	dsn := os.Getenv("DATABASE_URL")
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		//write migrations here
		db.AutoMigrate(&models.User{}, &models.Space{}, &models.Notes{})

		if err == nil {
			break
		}

		log.Println("Failed to connect to the database. Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Could not connect to the database after retries: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
