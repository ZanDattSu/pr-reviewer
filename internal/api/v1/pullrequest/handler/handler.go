package handler

import (
	api "github.com/ZanDattSu/pr-reviewer/internal/api/v1/pullrequest"
	"github.com/ZanDattSu/pr-reviewer/internal/service/pullrequest"
)

// Компиляторная проверка: убеждаемся, что *userService реализует интерфейс PRApi.
var _ api.PRApi = (*prHandler)(nil)

type prHandler struct {
	prService pullrequest.PRService
}

func NewPrHandler(prService pullrequest.PRService) *prHandler {
	return &prHandler{prService: prService}
}
