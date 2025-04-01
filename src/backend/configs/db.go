package configs

import (
	"fmt"
	"log"
	"os"

	"room-reservation/models" // ✅ import models

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	// ✅ ย้าย AutoMigrate มาที่นี่
	if err := database.AutoMigrate(
		&models.User{},
		// เพิ่ม model อื่น ๆ ที่นี่ได้ เช่น &models.Room{}, &models.Booking{}
	); err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	DB = database
	fmt.Println("✅ Connected to PostgreSQL & Migrated models!")
}
