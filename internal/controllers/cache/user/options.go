package user

import "context"

func (c *userCache) Check(ctx context.Context, key string) (bool, error) {
	// exists := c.cur.Exists(key)
	// exists.Result()
	return false, nil
}
