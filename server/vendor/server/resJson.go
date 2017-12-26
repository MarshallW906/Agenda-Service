package server

import (
	"net/http"

	"github.com/unrolled/render"
)

// OuterMsg : sturct of the response msg
type OuterMsg struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GenResponseMsg with the provided msg & data
func GenResponseMsg(messageCode int, data interface{}) *OuterMsg {
	message := http.StatusText(messageCode)
	return &OuterMsg{
		Msg:  message,
		Data: data,
	}
}

// ResponseOK ..
func ResponseOK(formatter *render.Render, w http.ResponseWriter, rawData interface{}) {
	formatter.JSON(w, http.StatusOK, GenResponseMsg(http.StatusOK, rawData))
}

// ResponseForbidden ..
func ResponseForbidden(formatter *render.Render, w http.ResponseWriter) {
	formatter.JSON(w, http.StatusForbidden, GenResponseMsg(http.StatusForbidden, struct{}{}))
}

// ResponseUnauthorized ..
func ResponseUnauthorized(formatter *render.Render, w http.ResponseWriter) {
	formatter.JSON(w, http.StatusUnauthorized, GenResponseMsg(http.StatusUnauthorized, struct{}{}))
}

// ResponseNotFound ..
func ResponseNotFound(formatter *render.Render, w http.ResponseWriter) {
	formatter.JSON(w, http.StatusNotFound, GenResponseMsg(http.StatusNotFound, struct{}{}))
}

// ResponseConflict ..
func ResponseConflict(formatter *render.Render, w http.ResponseWriter) {
	formatter.JSON(w, http.StatusConflict, GenResponseMsg(http.StatusConflict, struct{}{}))
}

// ResponseInternalServerError ..
func ResponseInternalServerError(formatter *render.Render, w http.ResponseWriter) {
	formatter.JSON(w, http.StatusInternalServerError, GenResponseMsg(http.StatusInternalServerError, struct{}{}))
}
