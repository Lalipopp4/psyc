package result

import (
	"context"
	"psyc/internal/models"
)

type Service interface {
	Get(ctx context.Context, key, param string) ([]models.Test, error)
	Keirsey(ctx context.Context, id string, res [4]int) error
	Hall(ctx context.Context, id string, res [5]int) error
	Bass(ctx context.Context, id string, self, task, social int) error
	Eysenck(ctx context.Context, id string, res int) error
}
