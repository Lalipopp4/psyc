package result

import (
	"context"
	"psyc/internal/models"
)

func (s *resultService) Get(ctx context.Context, key string, param string) ([]models.Test, error) {
	return []models.Test{}, nil
}

func (s *resultService) Keirsey(ctx context.Context, id string, res [4]int) error {
	return nil
}

func (s *resultService) Hall(ctx context.Context, id string, res [5]int) error {
	return nil
}

func (s *resultService) Bass(ctx context.Context, id string, self, task, social int) error {
	return nil
}

func (s *resultService) Eysenck(ctx context.Context, id string, res int) error {
	return nil
}
