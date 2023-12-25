package user

import (
	"context"
	"psyc/internal/models"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, string, error)
	Register(ctx context.Context, user *models.User) (string, error)
}
