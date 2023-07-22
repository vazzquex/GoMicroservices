package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
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
	var requestPlyload requestPayload

	err := app.readJSON(w, r, &requestPlyload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPlyload.Action {

	case "auth":
		app.authenticate(w, requestPlyload.Auth)

	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}
func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we will send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://auth-service:8080/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, errors.New("error creating new request: "+err.Error()))
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, errors.New("error making request to auth service: "+err.Error()))
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		// read the response body
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			app.errorJSON(w, errors.New("error reading response body: "+err.Error()))
			return
		}
		bodyString := string(bodyBytes)

		app.errorJSON(w, errors.New(fmt.Sprintf("error calling auth service, received status code %d: %s", response.StatusCode, bodyString)))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, errors.New("error decoding response from auth service: "+err.Error()))
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}
