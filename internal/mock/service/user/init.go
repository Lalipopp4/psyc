package user

import (
	"psyc/internal/controllers/db/user"

	cache "psyc/internal/controllers/cache"
)

type userService struct {
}

func New(repo user.Repository, cache cache.Cache) Service {
	return &userService{}
}
