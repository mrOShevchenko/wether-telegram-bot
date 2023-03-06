package internal

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"task2.3.3/internal/config"
	"task2.3.3/internal/logger"
)

type Container interface {
	NewConfig() *config.Config
	NewLogger() *zap.SugaredLogger
}

type container struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func NewContainer() (*container, error) {
	cfg, err := config.Get()
	if err != nil {
		return nil, errors.Wrap(err, "can't get configuration in container")
	}
	log, err := logger.Get()
	if err != nil {
		return nil, errors.Wrap(err, "can't get logger in container")
	}

	return &container{
		config: cfg,
		logger: log,
	}, nil
}

func (c *container) NewConfig() *config.Config {
	return c.config
}

func (c *container) NewLogger() *zap.SugaredLogger {
	return c.logger
}