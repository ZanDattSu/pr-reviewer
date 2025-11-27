package pullrequest

import (
	"github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *prRepository реализует интерфейс PullRequestRepository.
var _ PullRequestRepository = (*prRepository)(nil)

type prRepository struct {
	pool   *pgxpool.Pool
	getter *pgxv5.CtxGetter
}

func NewpPRRepository(pool *pgxpool.Pool) *prRepository {
	return &prRepository{
		pool:   pool,
		getter: pgxv5.DefaultCtxGetter,
	}
}
