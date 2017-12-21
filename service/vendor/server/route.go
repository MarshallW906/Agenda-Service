package server

import (
	"net/http"
	// Mnemonic
	_ "model"

	"github.com/gorilla/mux"
)

func router() http.Handler {
	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/token", tokenLoginHandler()).Methods("GET")
	theRouter.HandleFunc("/token", tokenLogoutHandler()).Methods("DELETE")
	theRouter.HandleFunc("/users/{username}", userRetriveHandler()).Methods("GET")
	theRouter.HandleFunc("/users/{username}", userDeleteHandler()).Methods("DELETE")
	theRouter.HandleFunc("/users", userCreateHandler()).Methods("PUT")
	theRouter.HandleFunc("/users", userListAllHandler()).Methods("GET")

	theRouter.Handle("/", theRouter.NotFoundHandler)
	return theRouter
}

// check if the request.username has logged in
// if so, simply return the corresponding token
// otherwise, check the username & password
//   if valid, generate a token based on the username & time and return
//   if not, return 403 with a "Forbidden" message
func tokenLoginHandler() http.HandlerFunc {
}

// check if the posted token is in the DB
// if so, simply remove it and return 200 with "OK" msg
// otherwise do nothing but return 401 with "Unauthorized"
func tokenLogoutHandler() http.HandlerFunc {

}

// First check if the client has logged in (token validation)
// if not, return 401 with "Unauthorized"
// if so, query the userinfo table and return 200(if found) or 404(not found)
func userRetriveHandler() http.HandlerFunc {
	//
}

// First check if the client has logged in (token validation)
// if not, return 401
// if so, query the username and
//        if exist, delete and return 200
// 	      otherwise return 404
func userDeleteHandler() http.HandlerFunc {

}

func userCreateHandler() http.HandlerFunc {

}

func userListAllHandler() http.HandlerFunc {

}
