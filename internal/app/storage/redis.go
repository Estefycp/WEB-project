package storage

import (
	"time"

	"../models"
	"github.com/go-redis/redis"
)

type redisStore struct {
	client *redis.Client
}

func (r *redisStore) SaveScore(player models.Player) error {
	return nil
}

func (r *redisStore) GetScores() ([]float64, error) {
	return nil, nil
}

func (r *redisStore) GetAliveDurations() ([]time.Duration, error) {
	return nil, nil
}
