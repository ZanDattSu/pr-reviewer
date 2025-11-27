package team

import (
	"github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *teamRepository реализует интерфейс TeamRepository.
var _ TeamRepository = (*teamRepository)(nil)

type teamRepository struct {
	pool   *pgxpool.Pool
	getter *pgxv5.CtxGetter
}

func NewTeamRepository(pool *pgxpool.Pool) *teamRepository {
	return &teamRepository{
		pool:   pool,
		getter: pgxv5.DefaultCtxGetter,
	}
}
