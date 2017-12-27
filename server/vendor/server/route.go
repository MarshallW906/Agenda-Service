package server

import (
	"entity"
	er "err"
	"logger"
	"model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var formatter *render.Render

func init() {
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
}

func router() http.Handler {
	theRouter := mux.NewRouter()

	theRouter.HandleFunc("/tokens", handlerTokenLogin()).Methods("POST")
	theRouter.HandleFunc("/tokens/{token_string}", handlerTokenLogout()).Methods("DELETE")
	theRouter.HandleFunc("/users/{username}", handlerUserRetrieve()).Methods("GET")
	theRouter.HandleFunc("/users/{username}", handlerUserDelete()).Methods("DELETE")
	theRouter.HandleFunc("/users", handlerUserCreate()).Methods("POST")
	theRouter.HandleFunc("/users", handlerUserListAll()).Methods("GET")

	theRouter.Handle("/", theRouter.NotFoundHandler)
	return theRouter
}

// check if the request.username has logged in
// if so, simply return the corresponding token
// otherwise, check the username & password
//   if valid, generate a token based on the username & time and return
//   if not, return 403 with a "Forbidden" message
func handlerTokenLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username := req.Form.Get("username")
		password := req.Form.Get("password")
		valid, err := model.CheckLoginInfo(username, password)
		if err != nil {
			logger.RecordErr(err)
			ResponseForbidden(formatter, w)
			return
		}
		if !valid {
			logger.RecordErr(er.WrongUsernameOrPassword)
			ResponseForbidden(formatter, w)
			return
		}
		token, err := model.GetToken(username)
		if err != nil {
			logger.FatalIf(err)
		}
		ResponseOK(formatter, w, &entity.Token{Token: token})
	}
}

// check if the posted token is in the DB
// if so, simply remove it and return 200 with "OK" msg
// otherwise do nothing but return 401 with "Unauthorized"
func handlerTokenLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		token := mux.Vars(req)["token_string"]
		if len(req.Form.Get("token")) != 0 {
			username, err := model.GetUsernameByToken(token)
			if err != nil {
				logger.RecordErr(err)
				ResponseUnauthorized(formatter, w)
				return
			}
			if len(username) == 0 {
				logger.RecordErr(er.RequireLoggedIn)
				ResponseUnauthorized(formatter, w)
				return
			}
			err = model.DeleteToken(token)
			if err != nil {
				logger.RecordErr(err)
				ResponseUnauthorized(formatter, w)
				return
			}
			ResponseOK(formatter, w, struct{}{})
			return
		}
		logger.RecordErr(er.RequireLoggedIn)
		ResponseUnauthorized(formatter, w)
	}
}

// First check if the client has logged in (token validation)
// if not, return 401 with "Unauthorized"
// if so, query the userinfo table and return 200(if found) or 404(not found)
func handlerUserRetrieve() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		token := req.Form.Get("token")
		valid, err := model.CheckTokenValid(token)
		if err != nil || !valid {
			logger.RecordErr(er.RequireLoggedIn)
			ResponseUnauthorized(formatter, w)
			return
		}
		username := mux.Vars(req)["username"]
		retUser, err := model.RetrieveUser(username)
		if err != nil {
			logger.RecordErr(err)
			ResponseNotFound(formatter, w)
			return
		}
		logger.Info("Retrieve user [%v]\n", username)
		ResponseOK(formatter, w, retUser)
	}
}

// First check if the client has logged in (token validation)
// if not, return 401
// if so, query the username and
//        if exist, delete and return 200
// 	      otherwise return 404
func handlerUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		token := req.Form.Get("token")
		valid, err := model.CheckTokenValid(token)
		if err != nil || !valid {
			logger.RecordErr(er.RequireLoggedIn)
			ResponseUnauthorized(formatter, w)
			return
		}
		username := mux.Vars(req)["username"]
		err = model.DeleteUser(username)
		if err != nil {
			logger.RecordErr(err)
			ResponseNotFound(formatter, w)
			return
		}
		err = model.DeleteTokenByUsername(username)
		if err != nil {
			logger.RecordErr(err)
			ResponseNotFound(formatter, w)
			return
		}
		ResponseOK(formatter, w, struct{}{})
	}
}

func handlerUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		username, password, email, phone :=
			req.Form.Get("username"), req.Form.Get("password"),
			req.Form.Get("email"), req.Form.Get("phone")
		newUser := &entity.User{
			Username: username,
			Password: password,
			Email:    email,
			Phone:    phone,
		}
		err := model.CreateUser(newUser)
		if err != nil {
			if err.Error() == er.UserAlreadyExists.Error() {
				logger.RecordErr(err)
				ResponseConflict(formatter, w)
			} else {
				ResponseInternalServerError(formatter, w)
				logger.FatalIf(err)
			}
			return
		}
		logger.Info("CreateUser [%+v]\n", username)
		ResponseCreated(formatter, w, newUser)
	}
}

func handlerUserListAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		token := req.Form.Get("token")
		valid, err := model.CheckTokenValid(token)
		if err != nil || !valid {
			logger.RecordErr(er.RequireLoggedIn)
			ResponseUnauthorized(formatter, w)
			return
		}
		username, err := model.GetUsernameByToken(token)
		if err != nil {
			logger.FatalIf(err)
			ResponseInternalServerError(formatter, w)
			return
		}
		allUsers, err := model.ListAllUsers()
		if err != nil {
			logger.FatalIf(err)
			ResponseInternalServerError(formatter, w)
			return
		}
		logger.Info("ListAllUsers By User:[%+v] with Token:[%+v]\n", username, token)
		ResponseOK(formatter, w, allUsers)
	}
}
