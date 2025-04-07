package exception

import "errors"

var (
	UserAlreadyExists            = errors.New("user already exists")
	InvalidEmail                 = errors.New("invalid email")
	InvalidPassword              = errors.New("invalid password")
	RefreshTokenIsAlreadyExpired = errors.New("refresh token is already expired")
)
