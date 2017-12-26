package model

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"

	"database"
	"entity"
	er "err"
)

// GetToken ..
func GetToken(username string) (string, error) {
	ret := &entity.Token{Username: username}
	has, err := database.Engine.Table("token").Get(ret)
	if err != nil {
		return "", err
	}
	if has == false {
		ret.Token, err = genToken(username)
		if err != nil {
			return "", err
		}
		_, err = database.Engine.Table("token").Insert(ret)
		if err != nil {
			return "", err
		}
	}
	return ret.Token, nil
}

// generate a token based on username & time, then return it
// no operations on database here
func genToken(username string) (string, error) {
	curtime := time.Now().Unix()
	h := md5.New()
	_, err := io.WriteString(h, strconv.FormatInt(curtime, 10)+username)
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token, nil
}

// DeleteToken ..
func DeleteToken(tokenStr string) error {
	t := &entity.Token{Token: tokenStr}
	has, err := database.Engine.Table("token").Get(t)
	if err != nil {
		return err
	}
	if !has {
		return er.RequireLoggedIn
	}
	_, err = database.Engine.Table("token").Delete(t)
	return err
}

// DeleteTokenByUsername ..
func DeleteTokenByUsername(username string) error {
	t := &entity.Token{Username: username}
	_, err := database.Engine.Table("token").Delete(t)
	return err
}

// CheckTokenValid : Token Valid if err == nil && BoolValue == true
func CheckTokenValid(tokenStr string) (bool, error) {
	return database.Engine.Table("token").Get(&entity.Token{Token: tokenStr})
}

// GetUsernameByToken ..
func GetUsernameByToken(tokenStr string) (string, error) {
	t := &entity.Token{Token: tokenStr}
	has, err := database.Engine.Table("token").Get(t)
	if err != nil {
		return "", err
	}
	if has {
		return t.Username, nil
	}
	return "", nil
}
