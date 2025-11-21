package config

import (
	"time"
)

type App interface {
	ShutdownTimeout() time.Duration
}

type ServerConfig interface {
	Address() string
	ResponseTimeout() time.Duration
	ReadHeaderTimeout() time.Duration
	ShutdownTimeout() time.Duration
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
	MigrationsPath() string
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
