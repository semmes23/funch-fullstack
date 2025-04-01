package routes

import (
	"room-reservation/controllers"
	"room-reservation/middleware"
	"room-reservation/repositories"
	"room-reservation/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRoutes เซ็ตอัพ routes พร้อม inject dependency
func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	authRoutes := router.Group("/auth")

	authRoutes.POST("/register", authController.Register)
	authRoutes.POST("/login", authController.Login)

	authRoutes.GET("/me", middleware.AuthMiddleware(), authController.Me)
}
