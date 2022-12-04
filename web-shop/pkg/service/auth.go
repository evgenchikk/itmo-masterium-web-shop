package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/evgenchikk/itmo-web-shop/models"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) SignUp(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.SignUp(user)
}

func generatePasswordHash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))

	return fmt.Sprintf("%s", hasher.Sum(nil))
}
