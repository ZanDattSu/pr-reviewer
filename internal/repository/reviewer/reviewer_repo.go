package reviewer

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type reviewerRepository struct {
	pool *pgxpool.Pool
}

func NewReviewerRepository(pool *pgxpool.Pool) *reviewerRepository {
	return &reviewerRepository{pool: pool}
}
