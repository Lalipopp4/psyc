package user

import redis "github.com/redis/go-redis/v9"

type userCache struct {
	cur *redis.Client
}

func New(cur *redis.Client) Cache {
	return &userCache{cur: cur}
}
