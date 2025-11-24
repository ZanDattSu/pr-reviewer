package handler

import (
	api "github.com/ZanDattSu/pr-reviewer/internal/api/v1/user"
	userService "github.com/ZanDattSu/pr-reviewer/internal/service/user"
)

// Компиляторная проверка: убеждаемся, что *userHandler реализует интерфейс UserApi.
var _ api.UserApi = (*userHandler)(nil)

type userHandler struct {
	userService userService.UserService
}

func NewUserHandler(userService userService.UserService) *userHandler {
	return &userHandler{userService: userService}
}
