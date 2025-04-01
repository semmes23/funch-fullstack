package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"room-reservation/dto"
	"room-reservation/models"
	"room-reservation/repositories"
	"room-reservation/utils"
)

type AuthService interface {
	RegisterUser(input dto.RegisterInput) error
	LoginUser(input dto.LoginInput) (string, error)
	GetProfile(userID uint) (*dto.UserProfileResponse, error)
}

type authService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) RegisterUser(input dto.RegisterInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// เปลี่ยนมาเรียกผ่าน repository
	return s.userRepo.Create(&user)
}

func (s *authService) LoginUser(input dto.LoginInput) (string, error) {
	// เปลี่ยนมาเรียกผ่าน repository
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("ไม่พบบัญชีผู้ใช้")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("รหัสผ่านไม่ถูกต้อง")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) GetProfile(userID uint) (*dto.UserProfileResponse, error) {
	// เปลี่ยนมาเรียกผ่าน repository
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	profile := &dto.UserProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return profile, nil
}
