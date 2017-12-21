package model

import (
	"database"
	"entity"
	er "err"
)

// CreateUser ..
func CreateUser(username, password, email, phone string) error {
	theUser := &entity.User{}
	has, err := database.Engine.Table("user").Where("user.username = ?", username).Get(theUser)
	if err != nil {
		return err
	}
	if has {
		return er.UserAlreadyExists
	}
	theUser = &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	_, err = database.Engine.Insert(theUser)
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
