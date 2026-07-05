package services

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/heronhoga/sme-marketing-kit-be/models"
	"github.com/heronhoga/sme-marketing-kit-be/repositories"
	"github.com/heronhoga/sme-marketing-kit-be/utils"
)

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) RegisterService(c context.Context, registerRequest models.RegisterRequest) error {
	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	registerRequest.Password = hashedPassword
	userId := uuid.NewString()
	if userId == "" {
		return errors.New("Internal Server Error")
	}

	err = s.repository.InsertNewUser(c, userId, registerRequest)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (s *AuthService) LoginService(c context.Context, loginRequest models.LoginRequest) (string, string, error) {
	currentUserPassword, currentUserId, err := s.repository.FindUserPasswordByEmail(c, loginRequest.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", errors.New("Invalid Credentials")
		}

		return "", "", errors.New("Internal Server Error")
	}

	if !utils.VerifyPassword(loginRequest.Password, currentUserPassword) {
		return "", "", errors.New("Invalid Credentials")
	}

	secret := os.Getenv("JWT_SECRET")
	accessToken, err := utils.GenerateAccessToken(currentUserId, loginRequest.Email, secret)
	if err != nil {
		return "", "", errors.New("Internal Server Error")
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return "", "", errors.New("Internal Server Error")
	}

	hashedRefreshToken := utils.HashRefreshToken(refreshToken)
	refreshTokenExpiredAt := time.Now().Add(24 * time.Hour)

	// save hashedRefreshToken
	refreshTokenId := uuid.NewString()
	err = s.repository.SaveRefreshToken(c, refreshTokenId, currentUserId, hashedRefreshToken, refreshTokenExpiredAt)
	if err != nil {
		return "", "", errors.New("Internal Server Error")
	}

	
	return accessToken, refreshToken, nil 
}
