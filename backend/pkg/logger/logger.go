package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defLogggerType = "prod"
)

type Config struct {
	Type string
}

func NewLogger(c ...Config) (*zap.SugaredLogger, error) {
	ltype := defLogggerType

	if len(c) > 0 {
		if c[0].Type != "" {
			ltype = c[0].Type
		}
	}

	var logger *zap.Logger
	var loggerConf zap.Config
	var err error

	switch ltype {
	case "dev":
		loggerConf = zap.NewDevelopmentConfig()
	case "prod":
		loggerConf = zap.NewProductionConfig()
	default:
		return nil, errors.New("unknown logger type")
	}

	loggerConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err = loggerConf.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
