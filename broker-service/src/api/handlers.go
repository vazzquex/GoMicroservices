package main

import (
	"errors"
	"net/http"
)

type requestPayLoad struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth, omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Request received from the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPlyload requestPayLoad

	err := app.readJSON(w, r, &requestPlyload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPlyload.Action {
	case "auth":

	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we will send to the auth microservice

	//call microservice

	//make sure we get back the correct status code
}
