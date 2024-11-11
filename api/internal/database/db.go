package database

import (
	"api/internal/config"
	"log"
    "fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initializes the database connection using GORM
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
