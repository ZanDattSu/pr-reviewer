package user

import (
	userService "github.com/ZanDattSu/pr-reviewer/internal/service/user"
)

// Компиляторная проверка: убеждаемся, что *userHandler реализует интерфейс UserApi.
var _ UserApi = (*userHandler)(nil)

type userHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *userHandler {
	return &userHandler{userService: userService}
}
