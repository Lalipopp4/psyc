package cache

import (
	"context"
)

func (c *userCache) Check(ctx context.Context, acl, key string) bool {
	return true
}

func (c *userCache) Add(ctx context.Context, key, value string) error {
	return nil
}
