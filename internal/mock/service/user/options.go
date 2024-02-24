package user

import (
	"context"
	"psyc/internal/models"
)

func (s *userService) Login(ctx context.Context, email, password string) (string, string, error) {
	return "", "/user", nil
}

func (s *userService) Register(ctx context.Context, user *models.User) (string, error) {
	return "", nil
}

func (s *userService) Update(ctx context.Context, info *models.User) error {
	return nil
}
