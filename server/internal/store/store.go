package store

import (
	redis "github.com/redis/go-redis/v9"
)

type Store struct {
	redisClient *redis.Client
}

func NewStore() *Store {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Store{
		redisClient: client,
	}
}
