package storage

import (
	"time"

	"../models"
	"github.com/go-redis/redis"
)

type redisStore struct {
	client *redis.Client
}

func (r *redisStore) SaveScore(player *models.Player) error {
	err := r.client.RPush("scores", player.Radius).Err()
	if err != nil {
		return err
	}
	t := time.Now()
	t.Sub(player.Born)
	err = r.client.RPush("alive", t.Sub(player.Born).String()).Err()
	return err
}

func (r *redisStore) GetScores() ([]float64, error) {
	return nil, nil
}

func (r *redisStore) GetAliveDurations() ([]time.Duration, error) {
	return nil, nil
}
