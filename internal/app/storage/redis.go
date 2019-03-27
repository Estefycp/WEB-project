package storage

import (
	"strconv"
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
	strs, err := r.client.LRange("scores", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var scores []float64
	for _, score := range strs {
		if n, err := strconv.ParseFloat(score, 64); err == nil {
			scores = append(scores, n)
		}
	}
	return scores, nil
}

func (r *redisStore) GetAliveDurations() ([]time.Duration, error) {
	strs, err := r.client.LRange("alive", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var durs []time.Duration
	for _, dur := range strs {
		if n, err := time.ParseDuration(dur); err == nil {
			durs = append(durs, n)
		}
	}
	return durs, nil
}
