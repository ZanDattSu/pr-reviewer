package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ZanDattSu/pr-reviewer/internal/config"
	reviewerV1 "github.com/ZanDattSu/pr-reviewer/shared/pkg/openapi/reviewer/v1"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(address string, api reviewerV1.Handler) (*HTTPServer, error) {
	openAPIHandler, err := reviewerV1.NewServer(api)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAPI server: %w", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(config.AppConfig().Server.ResponseTimeout()))

	r.Mount("/", openAPIHandler)

	server := &http.Server{
		Addr:              address,
		Handler:           r,
		ReadHeaderTimeout: config.AppConfig().Server.ReadHeaderTimeout(), // Защита от Slowloris атак:
		// тип DDoS-атаки, при которой атакующий умышленно медленно отправляет HTTP-заголовки,
		// удерживая соединения открытыми и истощая пул доступных соединений на сервере.
		// ReadHeaderTimeout принудительно закрывает соединение,
		// если клиент не успел отправить все заголовки за отведенное время.

	}

	return &HTTPServer{
		server: server,
	}, nil
}

func (s *HTTPServer) Serve() error {
	log.Println("Listening on", s.server.Addr)

	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("HTTP server shutdown error: %w", err)
	}

	return nil
}
