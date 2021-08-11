package mysql

import (
	"context"
	"fmt"
	"user-api/domain/models/user"
	user2 "user-api/domain/repository/user"
	"user-api/util"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterUser(ctx context.Context, user *user.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		if util.IsDuplicatedEntryError(err) {
			return fmt.Errorf("action=repo.registeruser user.email=%v err:%w", user.Email, user2.ErrDuplicateEntry)
		}
		return fmt.Errorf("action=repo.registeruser user.email=%v err:%v", user.Email, err)
	}
	return nil
}
