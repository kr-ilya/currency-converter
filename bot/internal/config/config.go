package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

const (
	ShutdownTimeout = 1 * time.Minute
)

type ServiceConfig struct {
	BotToken      string `env:"BOT_TOKEN,required"`
	WebhookBase   string `env:"WEBHOOK_BASE,required"`
	ListenAddress string `env:"LISTEN_ADDRESS,required"`
	LoggerType    string `env:"LOGGER_TYPE"`
	WebAppUrl     string `env:"WEBAPP_URL"`
}

func GetConfig() (*ServiceConfig, error) {
	var c ServiceConfig
	if err := envconfig.Process(context.Background(), &c); err != nil {
		return nil, err
	}
	return &c, nil
}
