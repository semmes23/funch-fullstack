package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"room-reservation/dto"
	"room-reservation/services"
)

// AuthController เป็น struct เพื่อใช้ inject dependency service เข้ามา
type AuthController struct {
	authService services.AuthService
}

// NewAuthController สร้าง controller ใหม่พร้อม inject service
func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register godoc
// @Summary ลงทะเบียนผู้ใช้
// @Description สร้างบัญชีใหม่ในระบบ
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterInput true "ข้อมูลผู้ใช้ใหม่"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.MessageResponse
// @Failure 500 {object} dto.MessageResponse
// @Router /auth/register [post]
func (ctrl *AuthController) Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.authService.RegisterUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ลงทะเบียนสำเร็จ"})
}

// Login godoc
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginInput true "Login Input"
// @Success 200 {object} dto.MessageResponse "Login Success"
// @Failure 400 {object} dto.MessageResponse "Invalid input"
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Me godoc
// @Summary Get User Profile
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserProfileResponse "User Profile"
// @Failure 401 {object} dto.MessageResponse "Unauthorized"
// @Failure 500 {object} dto.MessageResponse "Internal Server Error"
// @Router /auth/me [get]
// @Security BearerAuth
func (ctrl *AuthController) Me(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	profile, err := ctrl.authService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, profile)
}
