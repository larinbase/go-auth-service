package exception

import (
	"net/http"
)

type AppError struct {
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	UserAlreadyExists            = NewAppError("user already exists", http.StatusConflict)                  // 409
	InvalidEmail                 = NewAppError("invalid request", http.StatusBadRequest)                    // 400
	InvalidPassword              = NewAppError("invalid request", http.StatusBadRequest)                    // 400
	RefreshTokenIsAlreadyExpired = NewAppError("refresh token is already expired", http.StatusUnauthorized) //401
)
