package cache

import (
	"context"
	"psyc/internal/models"
)

type Cache interface {
	CheckUser(ctx context.Context, acl, key string) bool
	AddUser(ctx context.Context, key, value string) error
	GetReview(ctx context.Context, review *models.Review) (string, error)
	AddReview(ctx context.Context, review *models.Review) error
}
