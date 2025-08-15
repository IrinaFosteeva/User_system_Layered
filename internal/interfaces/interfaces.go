package interfaces

import (
	"context"
	"github.com/IrinaFosteeva/User_system_layered/internal/models"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	Create(ctx context.Context, u models.User) (models.User, error)
	Update(ctx context.Context, u models.User) (models.User, error)
	Delete(ctx context.Context, id int) error
}

type MainUserService interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id int) (models.User, error)
	Create(ctx context.Context, name, email string) (models.User, error)
	Update(ctx context.Context, id int, name, email string) (models.User, error)
	Delete(ctx context.Context, id int) error
}
