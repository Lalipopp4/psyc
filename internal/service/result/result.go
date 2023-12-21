package result

import (
	"context"
	"psyc/internal/models"
)

type Service interface {
	GetByUser(ctx context.Context, id string) ([]models.Result, error)
	Keirsey(ctx context.Context, id string, res [4]int) error
	Hall(ctx context.Context, id string, res [5]int) error
}
