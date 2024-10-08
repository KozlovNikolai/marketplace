package services

import (
	"context"
	"time"

	"github.com/KozlovNikolai/marketplace/internal/app/domain"
)

// UserService is a User service
type UserService struct {
	repo IUserRepository
}

// NewUserService creates a new User service
func NewUserService(repo IUserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s UserService) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	return s.repo.GetUserByLogin(ctx, login)
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	creatingTime := time.Now()
	var newUser = domain.NewUserData{
		Login:     user.Login(),
		Password:  user.Password(),
		Role:      "regular",
		Token:     "",
		CreatedAt: creatingTime,
		UpdatedAt: creatingTime,
	}
	creatingUser, err := domain.NewUser(newUser)
	if err != nil {
		return domain.User{}, err
	}
	return s.repo.CreateUser(ctx, creatingUser)
}

func (s UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repo.UpdateUser(ctx, user)
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s UserService) GetUsers(ctx context.Context, limit, offset int) ([]domain.User, error) {
	return s.repo.GetUsers(ctx, limit, offset)
}
