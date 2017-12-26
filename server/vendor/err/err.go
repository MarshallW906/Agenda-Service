package err

import (
	"errors"
	"logger"
	"strings"
)

var (
	UserAlreadyExists       error = errors.New("User already exists")
	WrongUsernameOrPassword error = errors.New("Wrong username or password")
	RequireLoggedIn         error = errors.New("Not logged in yet")
	UserNotExist            error = errors.New("User not found")
)

// RequireNonEmpty ..
func RequireNonEmpty(key string) error {
	return errors.New(strings.Title(key) + " must be a non-empty value")
}

// CheckErr ..
func CheckErr(err error) {
	if err != nil {
		logger.FatalIf(err)
	}
}
