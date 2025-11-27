package user

import (
	"github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *userRepository реализует интерфейс UserRepository.
var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	pool   *pgxpool.Pool
	getter *pgxv5.CtxGetter
}

func NewUserRepository(pool *pgxpool.Pool) *userRepository {
	return &userRepository{
		pool:   pool,
		getter: pgxv5.DefaultCtxGetter,
	}
}
