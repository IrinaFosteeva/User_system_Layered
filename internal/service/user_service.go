package service

import (
	"context"
	"errors"
	"github.com/IrinaFosteeva/User_system_layered/internal/models"
	"github.com/IrinaFosteeva/User_system_layered/internal/repository"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, name, email string) (models.User, error) {
	if name == "" || email == "" {
		return models.User{}, ErrInvalidInput
	}
	user := models.User{Name: name, Email: email}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, id int, name, email string) (models.User, error) {
	if name == "" || email == "" {
		return models.User{}, ErrInvalidInput
	}
	user := models.User{ID: id, Name: name, Email: email}
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
