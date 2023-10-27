package db

import (
	"colibri/pkg/db/models"
	"colibri/pkg/shared"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	db := GetDBInstance()
	log.Println("Migrating DB")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Community{})
	db.AutoMigrate(&models.Topic{})
	db.AutoMigrate(&models.Comments{})
	log.Println("✅  Migrated!")
}

func GetDBInstance() *gorm.DB {
	if db == nil {
		dbURL := shared.GetEnv("DATABASE_URL", "postgres://localhost:5432/colibri")

		log.Println("Connecting to DB")
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,         // Don't include params in the SQL log
				Colorful:                  true,          // Disable color
			},
		)
		newDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
			Logger: newLogger,
		})

		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		log.Println("✅  Connected!")
		db = newDB
	}

	return db
}
