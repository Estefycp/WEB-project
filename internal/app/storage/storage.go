package storage

import (
	"sync"
	"time"

	"../models"
	"github.com/go-redis/redis"
)

// Store interface for a Database
type Store interface {
	SaveScore(player models.Player) error
	GetScores() ([]float64, error)
	GetAliveDurations() ([]time.Duration, error)
}

var instance Store
var once sync.Once

// GetInstance to the DB connection (http://marcio.io/2015/07/singleton-pattern-in-go/)
func GetInstance() Store {
	once.Do(func() {
		instance = &redisStore{redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})}
	})
	return instance
}
