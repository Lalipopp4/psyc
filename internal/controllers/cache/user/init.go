package user

import "github.com/go-redis/redis"

type userCache struct {
	cur *redis.Conn
}

func New(cur *redis.Conn) Cache {
	return &userCache{cur: cur}
}
