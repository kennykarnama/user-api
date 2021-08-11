package mysql

import (
	"context"
	"gorm.io/gorm"
	"user-api/domain/models/user"
)

type repository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterUser(ctx context.Context, user *user.User) error {
	return nil
}
