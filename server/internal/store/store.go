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

func (s *Store) GetRedisClient() *redis.Client {
	return s.redisClient
}

func (s *Store) Close() error {
	if s.redisClient != nil {
		return s.redisClient.Close()
	}
	return nil
}
