package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/iqballm09/ezeats-backend-go/config"
	"github.com/iqballm09/ezeats-backend-go/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateModels(models []interface{}) {
	for _, model := range models {
		DB.AutoMigrate(model)
	}
}

func ConnectDB() {
	var err error

	// Get all configs for database
	host := config.Config("DB_HOST")
	username := config.Config("DB_USERNAME")
	password := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")
	port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)

	if err != nil {
		log.Println("Failed to parsing DB_PORT")
	}

	// create connection URL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	// connect to database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database!")
	}
	// create database pooling
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to create database pooling!")
	}
	// add configuration of database pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(120 * time.Minute)
	log.Println("Database has been connected!")

	// define models
	modelsData := []interface{}{
		&models.User{}, &models.Kategori{}, &models.Resto{},
		&models.Review{}, &models.Kabkota{}, &models.Kecamatan{},
		&models.Provinsi{}, &models.RestoDetail{}, &models.Koleksi{},
	}
	// auto migrate
	MigrateModels(modelsData)

	log.Println("Database migrated!")
}
