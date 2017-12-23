package err

import (
	"errors"
	"strings"
)

var (
	UserAlreadyExists       error = errors.New("User already exists")
	WrongUsernameOrPassword error = errors.New("Wrong username or passwordd")
	RequireLoggedIn         error = errors.New("Not logged in yet")
	UserNotExist            error = errors.New("User not found")
)

func RequireNonEmpty(key string) error {
	return errors.New(strings.Title(key) + " must be a non-empty value")
}
