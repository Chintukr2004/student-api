package service

import (
	"context"
	"errors"

	"github.com/Chintukr2004/student-api/internal/models"
	"github.com/Chintukr2004/student-api/internal/repository"
	"github.com/Chintukr2004/student-api/internal/utils"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Login(
	ctx context.Context,
	email, password, jwtSecret string,
) (string, error) {

	user, err := s.Repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID, "user", jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	if len(password) < 8 {
		return nil, errors.New("password too short")
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
	}

	err = s.Repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
