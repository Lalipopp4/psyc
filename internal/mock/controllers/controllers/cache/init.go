package cache

import redis "github.com/redis/go-redis/v9"

type userCache struct {
}

func New(cur *redis.Client) Cache {
	return &userCache{}
}
