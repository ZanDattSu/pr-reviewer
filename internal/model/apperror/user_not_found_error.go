package apperror

import "fmt"

type UserNotFoundError struct {
	userID string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user %s not found", e.userID)
}

func NewUserNotFoundError(userID string) *UserNotFoundError {
	return &UserNotFoundError{
		userID: userID,
	}
}

func (e *UserNotFoundError) UserID() string {
	return e.userID
}
