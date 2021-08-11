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

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var usr user.User
	err := r.db.Model(&user.User{}).Where("email = ?", email).Find(&usr).Error
	if err != nil {
		return nil, fmt.Errorf("action=repo.getUserByEmail email=%v err=%v", email, err)
	}
	// UUID empty will be considered as not exist user
	if usr.UUID == "" {
		return nil, nil
	}
	return &usr, nil
}

func (r *repository) GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	var usr user.User
	err := r.db.Model(&user.User{}).Where("id = ?", id).Find(&usr).Error
	if err != nil {
		return nil, fmt.Errorf("action=repo.getUserByID id=%v err=%v", id, err)
	}
	// UUID empty will be considered as not exist user
	if usr.UUID == "" {
		return nil, nil
	}
	return &usr, nil
}
