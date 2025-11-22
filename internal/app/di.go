package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	reviewerApi "github.com/ZanDattSu/pr-reviewer/internal/api/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/config"
	"github.com/ZanDattSu/pr-reviewer/pkg/closer"
)

type diContainer struct {
	api            reviewerApi.Api
	postgreSQLPool *pgxpool.Pool
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (di *diContainer) Api(ctx context.Context) reviewerApi.Api {
	if di.api == nil {
		// TODO Сделать DI
	}
	return di.api
}

func (di *diContainer) PostgreSQLPool(ctx context.Context) *pgxpool.Pool {
	if di.postgreSQLPool == nil {
		dbURI := config.AppConfig().Postgres.URI()

		pool, err := pgxpool.New(ctx, dbURI)
		if err != nil {
			panic(fmt.Sprintf("Failed to create pgxpool connect: %s", err))
		}

		err = pool.Ping(ctx)
		if err != nil {
			panic(fmt.Sprintf("Database is unavailable: %s", err))
		}

		closer.AddNamed("PostgreSQL pool", func(ctx context.Context) error {
			pool.Close()
			return nil
		})

		di.postgreSQLPool = pool
	}

	return di.postgreSQLPool
}
