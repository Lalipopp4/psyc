package result

import (
	"context"
	"psyc/internal/models"
)

type Repository interface {
	GetByTest(ctx context.Context, test string, params ...string) ([]models.Test, error)
	GetByUsers(ctx context.Context, users []string) ([]models.Test, error)
	Add(ctx context.Context, test *models.Test) error
	GetUsers(ctx context.Context, key string, param string) ([]string, error)
}
