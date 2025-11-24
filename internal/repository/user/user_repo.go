package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Компиляторная проверка: убеждаемся, что *userRepository реализует интерфейс UserRepository.
var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *userRepository {
	return &userRepository{pool: pool}
}
