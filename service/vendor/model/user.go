package model

import (
	"database"
	"entity"
	er "err"
)

// CreateUser ..
func CreateUser(newUser *entity.User) error {
	theUser := &entity.User{Username: newUser.Username}
	has, err := database.Engine.Table("user").Where("user.username = ?", theUser.Username).Get(theUser)
	if err != nil {
		return err
	}
	if has {
		return er.UserAlreadyExists
	}
	_, err = database.Engine.Insert(newUser)
	return err
}

// DeleteUser ..
func DeleteUser(username string) error {
	theUser := &entity.User{Username: username}
	has, err := database.Engine.Table("user").Get(theUser)
	if err != nil {
		return err
	}
	if has {
		_, err = database.Engine.Table("user").Delete(theUser)
	}
	return err
}

// RetriveUser ..
func RetriveUser(username string) (*entity.User, error) {
	theUser := &entity.User{Username: username}
	has, err := database.Engine.Table("user").Get(theUser)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, er.UserNotExist
	}
	return theUser, nil
}

// ListAllUsers ..
func ListAllUsers() (*entity.Users, error) {
	allUsers := new(entity.Users)
	err := database.Engine.Table("user").Find(allUsers)
	return allUsers, err
}

// CheckLoginInfo ..
func CheckLoginInfo(username, password string) (bool, error) {
	return database.Engine.Table("user").Get(&entity.User{
		Username: username,
		Password: password,
	})
}
