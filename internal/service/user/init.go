package user

import "psyc/internal/controllers/db/user"

type userService struct {
	repo user.Repository
}

func New(repo user.Repository) Service {
	return &userService{repo: repo}
}
