package user

import (
	"context"
	"psyc/internal/models"
)

func (r *userRepository) Add(ctx context.Context, user *models.User) error {
	return nil
}

func (r *userRepository) GetIDPassword(ctx context.Context, email string) (string, string) {
	return "", ""
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}
