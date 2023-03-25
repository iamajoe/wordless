package redis

import (
	"errors"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

type DB struct {
	db *redis.Client
}

func (db *DB) Close() error {
	return db.db.Close()
}

func Connect(host string, secret string, dbIndex uint) (*DB, error) {
	if dbIndex > 15 || dbIndex < 0 {
		return nil, errors.New(fmt.Sprintf("DB index needs to be between 0 and 15. Got %d", dbIndex))
	}

	db := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: secret,
		DB:       int(dbIndex),
	})

	return &DB{db}, nil
}
