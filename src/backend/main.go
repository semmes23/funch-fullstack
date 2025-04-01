// @title Room Reservation API
// @version 1.0
// @description Backend API for managing room reservation system.
// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter JWT Bearer token **_only_**
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"room-reservation/configs"
	"room-reservation/routes"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "room-reservation/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	configs.ConnectDB()
	db := configs.DB

	r := gin.Default()

	// ✅ Add Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ✅ เพิ่ม group "/api" ให้ตรงกับ Swagger BasePath
	api := r.Group("/api")
	routes.SetupAuthRoutes(api, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
