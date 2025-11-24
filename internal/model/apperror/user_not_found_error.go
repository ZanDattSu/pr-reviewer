package apperror

import "fmt"

type UserNotFoundError struct {
	userUUID string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user %s not found", e.userUUID)
}

func NewUserNotFoundError(name string) *UserNotFoundError {
	return &UserNotFoundError{
		userUUID: name,
	}
}
