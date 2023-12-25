package user

import (
	"context"
	"psyc/internal/models"
)

type Repository interface {
	Add(ctx context.Context, user *models.User) error
	GetIDPassword(ctx context.Context, email string) (string, string)
	Update(ctx context.Context, user *models.User) error
}
