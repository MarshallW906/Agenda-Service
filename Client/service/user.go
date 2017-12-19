package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ZM-J/Agenda-Service/Client/entity"
	"github.com/ZM-J/Agenda-Service/Client/err"
	"github.com/ZM-J/Agenda-Service/Client/logger"
	"github.com/ZM-J/Agenda-Service/Client/storage"
)

// Finish: Client
func Register(username string, password string, email string, phone string) {
	body := strings.NewReader(`{
		"username": ` + username + `,
		"password": ` + password + `,
		"email": ` + email + `,
		"phone": ` + phone + `
	  }`)
	req, err := http.NewRequest(http.MethodPut, "https://localhost:8080/users", body)
	logger.FatalIf(err)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Register successful.")
	} else {
		fmt.Println("Register failed.")
	}
}

// Finish: Client
func Login(username string, password string) {
	storage.RemoveSessionFile()
	resp, err := http.Get("https://localhost:8080/token?username=" + username + "&password=" + password)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Login successful.")
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		var data = inter["data"].(map[string]interface{})
		storage.StoreSession(&entity.Session{
			CurrentToken: data["token"].(string),
		})
	} else {
		fmt.Println("Login failed.")
	}
}

// Finish: Client
func Logout() {
	token, ok := storage.LoadCurToken()
	if !ok {
		return
	}
	req, err := http.NewRequest(http.MethodDelete, "https://localhost:8080/token?token="+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Logout successful.")
	} else {
		fmt.Println("Logout failed.")
	}
	storage.RemoveSessionFile()
}

// Finish: Client
func ListAllUsers() {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(err.RequireLoggedIn)
	}
	resp, err := http.Get("https://localhost:8080/users?token=" + token)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		data := inter["data"].(map[string]interface{})
		fmt.Printf("%-20s %-20s %-20s\n", "USERNAME", "EMAIL", "PHONE")
		for _, user := range data {
			user := user.(map[string]interface{})
			fmt.Printf("%-20s %-20s %-20s\n", user["username"], user["email"], user["phone"])
		}
	} else {
		fmt.Println("Unauthorized!")
	}
}

// Finish: Client
func RemoveUser(username string, password string) {
	token, ok := storage.LoadCurToken()
	if !ok {
		return
	}
	req, err := http.NewRequest(http.MethodDelete, "https://localhost:8080/users/?token="+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Delete successful.")
	} else {
		fmt.Println("Delete failed.")
	}
}
