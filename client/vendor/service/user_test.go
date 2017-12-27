package service

import (
	"fmt"
	"testing"
)

var (
	username string = "user"
	password string = "pass"
	email    string = "user@example.com"
	phone    string = "13333333333"
)

func TestRegister(t *testing.T) {
	fmt.Println("Testing Register...")
	Register(username, password, "mail", phone)
	t.Log("Test Register OK.")
}

func TestLogin(t *testing.T) {
	fmt.Println("Testing Login...")
	Login(username, password)
	t.Log("Test Login OK.")
}

// func TestLogout(t *testing.T) {
// 	fmt.Println("Testing Logout...")
// 	Login(username, password)
// 	// Logout()
// 	t.Log("Test Logout OK.")
// }

func TestListAllUsers(t *testing.T) {
	fmt.Println("Testing List All Users...")
	Login(username, password)
	ListAllUsers()
	t.Log("Test List All Users OK.")
}

func TestFindUser(t *testing.T) {
	fmt.Println("Testing Find User...")
	Login(username, password)
	FindUser(username)
	t.Log("Test Find User OK.")
}

func TestRemoveUser(t *testing.T) {
	fmt.Println("Testing Remove User...")
	Login(username, password)
	RemoveUser(username)
	t.Log("Test Remove User OK.")
}
