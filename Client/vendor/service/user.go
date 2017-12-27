package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"entity"
	e "err"
	"logger"
	"storage"
)

// Finish: Client
func Register(username string, password string, email string, phone string) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	form.Add("email", email)
	form.Add("phone", phone)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/users", strings.NewReader(form.Encode()))
	logger.FatalIf(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Register successful.")
	} else {
		fmt.Println("Register failed.")
		panic(e.UserAlreadyExists)
	}
}

// Finish: Client
func Login(username string, password string) {
	storage.RemoveSessionFile()
	resp, err := http.Get("http://localhost:8080/token?username=" + username + "&password=" + password)
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
		panic(e.WrongUsernameOrPassword)
	}
}

// Finish: Client
func Logout() {
	token, ok := storage.LoadCurToken()
	if !ok {
		return
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/token?token="+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Logout successful.")
	} else {
		fmt.Println("Logout failed.")
		panic(e.RequireLoggedIn)
	}
	storage.RemoveSessionFile()
}

// Finish: Client
func ListAllUsers() {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(e.RequireLoggedIn)
	}
	resp, err := http.Get("http://localhost:8080/users?token=" + token)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		data := inter["data"].([]interface{})
		fmt.Printf("%-20s %-20s %-20s\n", "USERNAME", "EMAIL", "PHONE")
		for _, user := range data {
			user := user.(map[string]interface{})
			fmt.Printf("%-20s %-20s %-20s\n", user["username"], user["email"], user["phone"])
		}
	} else {
		fmt.Println("Unauthorized!")
		panic(e.RequireLoggedIn)
	}
}

// Finish: Client
func RemoveUser(username string) {
	token, ok := storage.LoadCurToken()
	if !ok {
		return
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/"+username+"?token="+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Delete successful.")
	} else {
		fmt.Println("Delete failed.")
		panic(e.RequireLoggedIn)
	}
}

// TODO: Client
func FindUser(username string) {
	token, ok := storage.LoadCurToken()
	if !ok {
		return
	}
	resp, err := http.Get("http://localhost:8080/users/" + username + "?token=" + token)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		data := inter["data"].(map[string]interface{})
		fmt.Printf("%-20s %-20s %-20s\n", "USERNAME", "EMAIL", "PHONE")
		fmt.Printf("%-20s %-20s %-20s\n", data["username"], data["email"], data["phone"])
	} else if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("Unauthorized!")
		panic(e.RequireLoggedIn)
	} else {
		fmt.Println("User not exists.")
		panic(e.UserNotExist)
	}
}
