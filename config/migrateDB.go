package config

import (
	"log"

	"github.com/dawood-usman/go-ops/models"
)

func MigrateDB() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully!")
}
