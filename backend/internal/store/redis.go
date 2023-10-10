package store

import (
	"context"
	"currency-telegram-webapp-backend/internal/config"
	"currency-telegram-webapp-backend/internal/provider"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	*redis.Client
}

func NewRedisStore(c *config.RedisConfig) *RedisStore {
	return &RedisStore{
		redis.NewClient(&redis.Options{
			Addr:     c.Host + ":" + c.Port,
			Password: c.Pass,
			DB:       c.Db,
		}),
	}
}

func (r *RedisStore) SetRates(base provider.Currency, ts int64, rates *provider.Rates) error {
	ctx := context.Background()

	// set rates
	if _, err := r.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		for k, v := range *rates {
			r.HSet(ctx, "rates", string(k), v)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("set rates: %w", err)
	}

	// set base currency
	if err := r.Set(ctx, "base", string(base), 0).Err(); err != nil {
		return fmt.Errorf("set base currency: %w", err)
	}

	// set update rates timestamp
	if err := r.Set(ctx, "updateTime", strconv.FormatInt(ts, 10), 0).Err(); err != nil {
		return fmt.Errorf("set updateTime: %w", err)
	}

	return nil
}

func (r *RedisStore) GetRates() (*provider.Rates, error) {
	ctx := context.Background()

	var rates = make(provider.Rates)
	res, err := r.HGetAll(ctx, "rates").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if errors.Is(err, redis.Nil) {
		return nil, ErrNotFound
	}

	for k, v := range res {
		if r, e := strconv.ParseFloat(v, 64); e == nil {
			rates[provider.Currency(k)] = r
		} else {
			return nil, e
		}
	}

	return &rates, nil
}

func (r *RedisStore) GetUpdateTime() (int64, error) {
	ctx := context.Background()

	var updateTime int64
	err := r.Get(ctx, "updateTime").Scan(&updateTime)
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	if errors.Is(err, redis.Nil) {
		return 0, ErrNotFound
	}

	return updateTime, nil
}

func (r *RedisStore) GetBase() (provider.Currency, error) {
	ctx := context.Background()

	var base string
	err := r.Get(ctx, "base").Scan(&base)
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}

	if errors.Is(err, redis.Nil) {
		return "", ErrNotFound
	}

	return provider.Currency(base), nil
}
