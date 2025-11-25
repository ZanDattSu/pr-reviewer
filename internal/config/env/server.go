package env

import (
	"net"
	"time"

	"github.com/caarlos0/env/v11"
)

type serverEnvConfig struct {
	HTTPHost              string        `env:"HTTP_HOST,required"`
	HTTPPort              string        `env:"HTTP_PORT,required"`
	HTTPResponseTimeout   time.Duration `env:"HTTP_RESPONSE_TIMEOUT,required"`
	HttpReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT,required"`
	HttpShutdownTimeout   time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT,required"`
}

type serverConfig struct {
	raw serverEnvConfig
}

func NewServerConfig() (*serverConfig, error) {
	var raw serverEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &serverConfig{raw: raw}, nil
}

func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.raw.HTTPHost, cfg.raw.HTTPPort)
}

func (cfg *serverConfig) ResponseTimeout() time.Duration {
	return cfg.raw.HTTPResponseTimeout
}

func (cfg *serverConfig) ReadHeaderTimeout() time.Duration {
	return cfg.raw.HttpReadHeaderTimeout
}

func (cfg *serverConfig) ShutdownTimeout() time.Duration {
	return cfg.raw.HttpShutdownTimeout
}
