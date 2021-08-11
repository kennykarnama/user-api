package user

import (
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	userEntity "user-api/domain/models/user"
	userRepo "user-api/domain/repository/user"
	"user-api/util"
)

type service struct {
	repo userRepo.Repository
}

func NewService(repo userRepo.Repository) *service {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, user *userEntity.User) error {
	user.UUID = uuid.NewV4().String()
	hashedPwd, err := util.Encrypt(user.PasswordAsByte())
	if err != nil {
		return err
	}
	user.Password = hashedPwd
	err = s.repo.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, userRepo.ErrDuplicateEntry) {
			return ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*userEntity.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *service) GetUserByID(ctx context.Context, id int64) (*userEntity.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
