package reviewer

import (
	"github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type reviewerRepository struct {
	pool   *pgxpool.Pool
	getter *pgxv5.CtxGetter
}

func NewReviewerRepository(pool *pgxpool.Pool) *reviewerRepository {
	return &reviewerRepository{
		pool:   pool,
		getter: pgxv5.DefaultCtxGetter,
	}
}
