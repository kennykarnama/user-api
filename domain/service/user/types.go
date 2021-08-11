package user

import (
	"context"
	"user-api/domain/models/user"
)

type Service interface {
	RegisterUser(ctx context.Context, newUser *user.User) error
}
