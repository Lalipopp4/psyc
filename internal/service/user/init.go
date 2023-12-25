package user

import (
	"psyc/internal/controllers/db/user"

	cache "psyc/internal/controllers/cache"
)

type userService struct {
	repo  user.Repository
	cache cache.Cache
}

func New(repo user.Repository, cache cache.Cache) Service {
	return &userService{repo: repo}
}
