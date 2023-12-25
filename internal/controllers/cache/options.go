package user

import (
	"context"
)

func (c *userCache) Check(ctx context.Context, acl, key string) bool {
	return c.cur.HExists(ctx, acl, key).Val()
}

func (c *userCache) Add(ctx context.Context, key, value string) error {
	return c.cur.HSet(ctx, "user", key, value).Err()
}
