package store

import (
	"currency-telegram-webapp-backend/internal/provider"
	"errors"
)

type Store interface {
	SetRates(base provider.Currency, ts int64, rates *provider.Rates) error
	GetUpdateTime() (int64, error)
	GetRates() (*provider.Rates, error)
	GetBase() (provider.Currency, error)
}

var ErrNotFound = errors.New("not found")
