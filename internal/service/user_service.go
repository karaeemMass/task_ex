package service

import (
	"context"

	"task_ex/internal/auth"
	"task_ex/internal/model"
	"task_ex/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo       *repository.UserRepository
	jwtManager *auth.JWT
}

func NewUserService(repo *repository.UserRepository, jwtManager *auth.JWT) *UserService {
	return &UserService{repo: repo, jwtManager: jwtManager}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*model.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.ListUsers(ctx)
}
