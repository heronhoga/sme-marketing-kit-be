package services

import (
	"context"
	"errors"

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