package pullrequest

import (
	"github.com/ZanDattSu/pr-reviewer/internal/service/pullrequest"
)

// Компиляторная проверка: убеждаемся, что *userService реализует интерфейс PRApi.
var _ PRApi = (*prHandler)(nil)

type prHandler struct {
	prService pullrequest.PRService
}

func NewPrHandler(prService pullrequest.PRService) *prHandler {
	return &prHandler{prService: prService}
}
