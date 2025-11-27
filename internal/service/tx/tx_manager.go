package tx

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"
)

// TransactionManager повторяет интерфейс avito-tech/go-transaction-manager/trm Manager для генерации моков через Mockery
type TransactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
	DoWithSettings(ctx context.Context, s trm.Settings, fn func(ctx context.Context) error) error
}
