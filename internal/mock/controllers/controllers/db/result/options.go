package result

import (
	"context"
	"psyc/internal/models"
)

func (r *resultRepository) GetUsers(ctx context.Context, key, param string) ([]string, error) {
	return []string{}, nil
}

func (r *resultRepository) GetByTest(ctx context.Context, test string, params ...string) ([]models.Test, error) {
	return []models.Test{}, nil
}

func (r *resultRepository) GetByUsers(ctx context.Context, users []string) ([]models.Test, error) {
	return []models.Test{}, nil

}

func (r *resultRepository) Add(ctx context.Context, test *models.Test) error {
	return nil
}
