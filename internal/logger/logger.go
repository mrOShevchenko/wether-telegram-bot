package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func Get() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("can't create logger zap")
	}
	defer logger.Sync()
	log := logger.Sugar()
	return log, nil
}
