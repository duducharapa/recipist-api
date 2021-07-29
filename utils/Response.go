package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int
	Json       interface{}
}

type HttpError struct {
	Error string
}

func send(w http.ResponseWriter, res *Response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(res.Json)
	return
}

func internalError(w http.ResponseWriter) {
	w.WriteHeader(500)
	return
}

func Send(w http.ResponseWriter, json interface{}, code int) {
	send(w, &Response{Json: json, StatusCode: code})
	return
}

func SendError(w http.ResponseWriter, msg string, code int) {
	jsonErr, err := json.Marshal(&HttpError{Error: msg})

	if err != nil {
		internalError(w)
		return
	}

	send(w, &Response{Json: jsonErr, StatusCode: code})
	return
}
