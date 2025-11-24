package team

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *teamRepository реализует интерфейс TeamRepository.
var _ TeamRepository = (*teamRepository)(nil)

type teamRepository struct {
	pool *pgxpool.Pool
}

func NewTeamRepository(pool *pgxpool.Pool) *teamRepository {
	return &teamRepository{pool: pool}
}
