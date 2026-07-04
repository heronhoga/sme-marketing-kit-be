package services

import "github.com/heronhoga/sme-marketing-kit-be/repositories"

type AuthService struct {
	repository *repositories.AuthRepository
}

func NewAuthService(repository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}