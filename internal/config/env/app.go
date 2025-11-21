package env

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type appEnvConfig struct {
	ShutTimeout time.Duration `env:"SHUTDOWN_TIMEOUT,required"`
}

type appConfig struct {
	raw appEnvConfig
}

func NewAppConfig() (*appConfig, error) {
	var raw appEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &appConfig{raw: raw}, nil
}

func (a *appConfig) ShutdownTimeout() time.Duration {
	return a.raw.ShutTimeout
}
