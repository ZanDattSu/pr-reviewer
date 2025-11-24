package pullrequest

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *prRepository реализует интерфейс PullRequestRepository.
var _ PullRequestRepository = (*prRepository)(nil)

type prRepository struct {
	pool *pgxpool.Pool
}

func NewpPRRepository(pool *pgxpool.Pool) *prRepository {
	return &prRepository{pool: pool}
}
