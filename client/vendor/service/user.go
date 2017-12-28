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

var (
	localhostURI  = "http://localhost:8080"
	mockserverURI = "https://private-cf3af9-agenda33.apiary-mock.com"
)

func Register(username string, password string, email string, phone string) {
	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)
	form.Add("email", email)
	form.Add("phone", phone)

	req, err := http.NewRequest(http.MethodPost, mockserverURI+"/users", strings.NewReader(form.Encode()))
	logger.FatalIf(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		logger.Info("Register successful.")
	} else {
		logger.FatalIf(e.UserAlreadyExists)
	}
}

func Login(username string, password string) {
	storage.RemoveSessionFile()

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, mockserverURI+"/tokens", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		logger.Info("Login successful.")
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		var data = inter["data"].(map[string]interface{})
		storage.StoreSession(&entity.Session{
			CurrentToken: data["token"].(string),
		})
	} else {
		logger.FatalIf(e.WrongUsernameOrPassword)
	}
}

func Logout() {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(e.RequireLoggedIn)
	}
	req, err := http.NewRequest(http.MethodDelete, mockserverURI+"/tokens/"+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		logger.Info("Logout successful.")
	} else {
		logger.FatalIf(e.RequireLoggedIn)
	}
	storage.RemoveSessionFile()
}

func ListAllUsers() {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(e.RequireLoggedIn)
	}
	resp, err := http.Get(mockserverURI + "/users?token=" + token)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("ok 1")
		// _, err := ioutil.ReadAll(resp.Body)
		// logger.FatalIf(err)
		// fmt.Println(body)
		fmt.Println("ok 1.5")
		var inter map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&inter)
		logger.FatalIf(err)
		fmt.Println("ok 2")
		data := inter["data"].([]interface{})
		fmt.Printf("%-20s %-20s %-20s\n", "USERNAME", "EMAIL", "PHONE")
		for _, user := range data {
			user := user.(map[string]interface{})
			fmt.Printf("%-20s %-20s %-20s\n", user["username"], user["email"], user["phone"])
		}
	} else {
		logger.FatalIf(e.RequireLoggedIn)
	}
}

func RemoveUser(username string) {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(e.RequireLoggedIn)
	}
	req, err := http.NewRequest(http.MethodDelete, mockserverURI+"/users/"+username+"?token="+token, nil)
	logger.FatalIf(err)
	resp, err := http.DefaultClient.Do(req)
	logger.FatalIf(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		logger.Info("Delete successful.")
	} else {
		logger.FatalIf(e.RequireLoggedIn)
	}
}

func FindUser(username string) {
	token, ok := storage.LoadCurToken()
	if !ok {
		logger.FatalIf(e.RequireLoggedIn)
	}
	resp, err := http.Get(mockserverURI + "/users/" + username + "?token=" + token)
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
		logger.FatalIf(e.RequireLoggedIn)
	} else {
		logger.FatalIf(e.UserNotExist)
	}
}
