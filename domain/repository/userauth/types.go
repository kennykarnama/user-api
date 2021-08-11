package userauth

import (
	"context"
	"user-api/domain/models/auth"
)

type Repository interface {
	SaveJWTSession(ctx context.Context, userID int64, metadata *auth.JWTMetadata) error
	IsAccessTokenExist(ctx context.Context, token string) (bool, error)
	IsRefreshTokenExist(ctx context.Context, token string) (bool, error)
	DeleteJWTSession(ctx context.Context, token string) error
}
