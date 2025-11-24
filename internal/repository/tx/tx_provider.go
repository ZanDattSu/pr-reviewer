package tx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Provider interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

type provider struct {
	pool *pgxpool.Pool
}

func NewTxProvider(pool *pgxpool.Pool) *provider {
	return &provider{pool: pool}
}

func (p *provider) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, opts)
}
