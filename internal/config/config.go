package config

import (
	"github.com/joho/godotenv"

	"github.com/ZanDattSu/pr-reviewer/internal/config/env"
)

var appConfig *config

type config struct {
	App      App
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil {
		return err
	}

	app, err := env.NewAppConfig()
	if err != nil {
		return err
	}

	order, err := env.NewServerConfig()
	if err != nil {
		return err
	}

	postgres, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	logger, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		App:      app,
		Server:   order,
		Postgres: postgres,
		Logger:   logger,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
