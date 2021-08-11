package userauth

import (
	"context"
	"user-api/domain/models/auth"
	userEntity "user-api/domain/models/user"
)

type LoginRequest struct {
	Email      string
	Password   string
	DeviceID   string
	DeviceName string
}

type ValidateTokenResult struct {
	User       *userEntity.User
	DeviceID   string
	DeviceName string
}

type Service interface {
	Login(ctx context.Context, req LoginRequest) (*auth.JWTMetadata, error)
	ValidateToken(ctx context.Context, token string) (*ValidateTokenResult, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (*auth.JWTMetadata, error)
}
