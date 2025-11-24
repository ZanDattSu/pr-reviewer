package user

import (
	"github.com/ZanDattSu/pr-reviewer/internal/repository/user"
)

// Компиляторная проверка: убеждаемся, что *userService реализует интерфейс UserService.
var _ UserService = (*userService)(nil)

type userService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}
