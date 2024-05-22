package cache

import (
	"context"
	"psyc/internal/models"
)

func (c *userCache) CheckUser(ctx context.Context, acl, key string) bool {
	return c.cur.HExists(ctx, acl, key).Val()
}

func (c *userCache) AddUser(ctx context.Context, key, value string) error {
	return c.cur.Set(ctx, key, value, c.cfg.TTLToken).Err()
}

func (c *userCache) AddReview(ctx context.Context, review *models.Review) error {
	return c.cur.Set(ctx, review.ResultID,
		review.Review, c.cfg.TTLData).Err()
}

func (c *userCache) GetReview(ctx context.Context, review *models.Review) (string, error) {
	return c.cur.Get(ctx, review.ResultID).Result()
}
