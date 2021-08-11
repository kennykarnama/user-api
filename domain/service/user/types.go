package user

import (
	"context"
	"user-api/domain/models/user"
)

type Service interface {
	RegisterUser(ctx context.Context, newUser *user.User) error
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
	GetUserByID(ctx context.Context, id int64) (*user.User, error)
}
