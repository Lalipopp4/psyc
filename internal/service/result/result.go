package result

import (
	"context"
	"psyc/internal/models"
)

type Service interface {
	GetTest(ctx context.Context, test *models.Test) (map[uint]*models.Question, error)
	AddResult(ctx context.Context, result *models.Test) (string, error)
	AddReview(ctx context.Context, result *models.Review) error
	AddTest(ctx context.Context, test *models.Test) error
	GetReview(ctx context.Context, review *models.Review) (string, error)
	GetResults(ctx context.Context, user *models.User) ([]models.Result, error)
}
