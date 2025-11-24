package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	reviewerApi "github.com/ZanDattSu/pr-reviewer/internal/api/v1"
	"github.com/ZanDattSu/pr-reviewer/internal/api/v1/health"
	prApi "github.com/ZanDattSu/pr-reviewer/internal/api/v1/pullrequest"
	prHandler "github.com/ZanDattSu/pr-reviewer/internal/api/v1/pullrequest/handler"
	teamApi "github.com/ZanDattSu/pr-reviewer/internal/api/v1/team"
	userApi "github.com/ZanDattSu/pr-reviewer/internal/api/v1/user"
	userHandler "github.com/ZanDattSu/pr-reviewer/internal/api/v1/user/handler"
	"github.com/ZanDattSu/pr-reviewer/internal/config"
	prRepository "github.com/ZanDattSu/pr-reviewer/internal/repository/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/reviewer"
	teamRepository "github.com/ZanDattSu/pr-reviewer/internal/repository/team"
	"github.com/ZanDattSu/pr-reviewer/internal/repository/tx"
	userRepository "github.com/ZanDattSu/pr-reviewer/internal/repository/user"
	prService "github.com/ZanDattSu/pr-reviewer/internal/service/pullrequest"
	teamService "github.com/ZanDattSu/pr-reviewer/internal/service/team"
	userService "github.com/ZanDattSu/pr-reviewer/internal/service/user"
	"github.com/ZanDattSu/pr-reviewer/pkg/closer"
)

type diContainer struct {
	api reviewerApi.Api

	healthApi health.HealthApi
	prApi     prApi.PRApi
	teamApi   teamApi.TeamApi
	userApi   userApi.UserApi

	prService   prService.PRService
	teamService teamService.TeamService
	userService userService.UserService

	prRepository       prRepository.PullRequestRepository
	reviewerRepository reviewer.ReviewerRepository
	teamRepository     teamRepository.TeamRepository
	userRepository     userRepository.UserRepository

	tx             tx.Provider
	postgreSQLPool *pgxpool.Pool
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

// API

func (di *diContainer) Api(ctx context.Context) reviewerApi.Api {
	if di.api == nil {
		di.api = reviewerApi.NewApi(
			di.HealthApi(ctx),
			di.PRApi(ctx),
			di.TeamApi(ctx),
			di.UserApi(ctx),
		)
	}
	return di.api
}

func (di *diContainer) HealthApi(ctx context.Context) health.HealthApi {
	if di.healthApi == nil {
		di.healthApi = health.NewHealthHandler()
	}
	return di.healthApi
}

func (di *diContainer) PRApi(ctx context.Context) prApi.PRApi {
	if di.prApi == nil {
		di.prApi = prHandler.NewPrHandler(di.PRService(ctx))
	}
	return di.prApi
}

func (di *diContainer) TeamApi(ctx context.Context) teamApi.TeamApi {
	if di.teamApi == nil {
		di.teamApi = teamApi.NewTeamHandler(di.TeamService(ctx))
	}
	return di.teamApi
}

func (di *diContainer) UserApi(ctx context.Context) userApi.UserApi {
	if di.userApi == nil {
		di.userApi = userHandler.NewUserHandler(di.UserService(ctx))
	}
	return di.userApi
}

// SERVICE

func (di *diContainer) PRService(ctx context.Context) prService.PRService {
	if di.prService == nil {
		di.prService = prService.NewPrService(
			di.PRRepository(ctx),
			di.ReviewerRepository(ctx),
			di.TeamRepository(ctx),
			di.TxProvider(ctx),
		)
	}
	return di.prService
}

func (di *diContainer) TeamService(ctx context.Context) teamService.TeamService {
	if di.teamService == nil {
		di.teamService = teamService.NewTeamService(di.TeamRepository(ctx))
	}
	return di.teamService
}

func (di *diContainer) UserService(ctx context.Context) userService.UserService {
	if di.userService == nil {
		di.userService = userService.NewUserService(di.UserRepository(ctx))
	}
	return di.userService
}

// REPOSITORY

func (di *diContainer) PRRepository(ctx context.Context) prRepository.PullRequestRepository {
	if di.prRepository == nil {
		di.prRepository = prRepository.NewpPRRepository(di.PostgreSQLPool(ctx))
	}
	return di.prRepository
}

func (di *diContainer) ReviewerRepository(ctx context.Context) reviewer.ReviewerRepository {
	if di.reviewerRepository == nil {
		di.reviewerRepository = reviewer.NewReviewerRepository(di.PostgreSQLPool(ctx))
	}
	return di.reviewerRepository
}

func (di *diContainer) TeamRepository(ctx context.Context) teamRepository.TeamRepository {
	if di.teamRepository == nil {
		di.teamRepository = teamRepository.NewTeamRepository(di.PostgreSQLPool(ctx))
	}
	return di.teamRepository
}

func (di *diContainer) UserRepository(ctx context.Context) userRepository.UserRepository {
	if di.userRepository == nil {
		di.userRepository = userRepository.NewUserRepository(di.PostgreSQLPool(ctx))
	}
	return di.userRepository
}

func (di *diContainer) TxProvider(ctx context.Context) tx.Provider {
	if di.tx == nil {
		di.tx = tx.NewTxProvider(di.PostgreSQLPool(ctx))
	}
	return di.tx
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
