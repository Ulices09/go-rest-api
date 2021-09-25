package logger

import (
	"go-rest-api/internal/config"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func New(config config.Config) (logger Logger, err error) {
	var zapLogger *zap.Logger

	if config.IsProduction() {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		return
	}

	logger.SugaredLogger = zapLogger.Sugar()

	return
}
