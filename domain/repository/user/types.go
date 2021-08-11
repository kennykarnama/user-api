package user

import (
	"context"
	"user-api/domain/models/user"
)

type Repository interface {
	RegisterUser(ctx context.Context, user *user.User) error
}
