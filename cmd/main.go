package main

import (
	"fmt"

	apiV1 "github.com/ZanDattSu/pr-reviewer/internal/api/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/config"
	"github.com/ZanDattSu/pr-reviewer/internal/server"
)

const configPath = ".env"

func main() {
	if err := config.Load(configPath); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	api := apiV1.NewApi()

	httpServer, err := server.NewHTTPServer(config.AppConfig().Server.Address(), api)
	if err != nil {
		return
	}

	err = httpServer.Serve()
	if err != nil {
		return
	}
}
