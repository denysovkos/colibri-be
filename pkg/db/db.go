package db

import (
	"colibri/pkg/db/models"
	"colibri/pkg/shared"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db := GetDBInstance()
	log.Println("Migrating DB")
	db.AutoMigrate(&models.User{})
	log.Println("✅  Migrated!")
}

func GetDBInstance() *gorm.DB {
	if db == nil {
		dbURL := shared.GetEnv("DATABASE_URL", "postgres://localhost:5432/colibri")

		log.Println("Connecting to DB")
		newDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		log.Println("✅  Connected!")
		db = newDB
	}

	return db
}
