package cache

import redis "github.com/redis/go-redis/v9"

type userCache struct {
	cur *redis.Client
	cfg *Config
}

func New(cur *redis.Client, cfg *Config) Cache {
	return &userCache{cur: cur}
}
