package result

import (
	"context"
	"psyc/internal/models"
)

type Repository interface {
	Get(ctx context.Context, param interface{}) ([]models.Result, error)
	Add(ctx context.Context, result models.Result) error
}
