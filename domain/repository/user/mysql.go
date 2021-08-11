package user

import (
	"context"
	"gorm.io/gorm"
	"user-api/domain/models/user"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *mysqlRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) RegisterUser(ctx context.Context, user *user.User) error {
	return nil
}
