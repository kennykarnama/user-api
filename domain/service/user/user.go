package user

import "user-api/domain/repository/user"

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) *service {
	return &service{repo: repo}
}
