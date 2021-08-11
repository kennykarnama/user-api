package user

import (
	"context"
	uuid "github.com/satori/go.uuid"
	userEntity "user-api/domain/models/user"
	userRepo "user-api/domain/repository/user"
)

type service struct {
	repo userRepo.Repository
}

func NewService(repo userRepo.Repository) *service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, user *userEntity.User) error {
	user.UUID = uuid.NewV4().String()
	return s.repo.RegisterUser(ctx, user)
}
