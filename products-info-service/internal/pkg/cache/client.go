package cache

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(address, password string) *Redis {
	return &Redis{
		redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       0,
		}),
	}
}
