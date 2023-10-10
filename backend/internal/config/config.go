package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

const (
	RedisExpire     = 1 * time.Hour
	ShutdownTimeout = 1 * time.Minute
	UpdatePeriod    = 60 * 60 // 1 hour

	DefaultFromCurrency = "EUR"
	DefaultToCurrency   = "USD"

	// Maximum number of symbols per request
	ReqMaxSymbols = 5
)

type ServiceConfig struct {
	Redis            RedisConfig
	ListenAddress    string `env:"LISTEN_ADDRESS,required"`
	LoggerType       string `env:"LOGGER_TYPE"`
	FixerAccessToken string `env:"FIXER_ACCESS_TOKEN,required"`
	BotToken         string `env:"BOT_TOKEN,required"`
}

type RedisConfig struct {
	Db   int    `env:"REDIS_DB,required"`
	Pass string `env:"REDIS_PASS,required"`
	Host string `env:"REDIS_HOST,required"`
	Port string `env:"REDIS_PORT,required"`
}

func GetConfig() (*ServiceConfig, error) {
	var c ServiceConfig
	if err := envconfig.Process(context.Background(), &c); err != nil {
		return nil, err
	}
	return &c, nil
}
