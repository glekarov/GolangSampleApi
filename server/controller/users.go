package controller

import (
	"SampleApiServerGL/server/auth"
	"SampleApiServerGL/server/jsons"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	invalidMethod        = "The requested method is not valid"
	invalidRequest       = "Can not process the request sent"
	incorrectCredentials = "Incorrect Username or Password"
	userAlreadyExists    = "User already exists"
)

// Creates the controller to process users
func CreateUsersControllers(router *mux.Router) {
	router.HandleFunc("/user/validate", validateUserCredentials)
	router.HandleFunc("/user/add", createNewUser)
}

// Prepare a JSON response with the appropriate error code
func writeErrorResponse(w http.ResponseWriter, code int, msg string) {
	resp := jsons.NewResponse(code, msg)
	fmt.Fprintf(w, "%v", resp)
	w.WriteHeader(code)
}

// Validates whether the credentials passed to the service are valid
// returns an error messages correspongind to the situation
func validateUserCredentials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Expected POST methcd but found: %v", r.Method)
		writeErrorResponse(w, http.StatusMethodNotAllowed, invalidMethod)
		return
	}

	var jsonUser jsons.JsonUserPassword
	reqBody, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal([]byte(reqBody), &jsonUser); err != nil {
		log.Printf("Can not parse request, %v", err)
		writeErrorResponse(w, http.StatusBadRequest, invalidRequest)
		return
	}

	err := auth.ValidateUserAndPassword(jsonUser.Username, jsonUser.Password)
	if err != nil {
		log.Printf("Can not parse request, %v", err)
		writeErrorResponse(w, http.StatusForbidden, incorrectCredentials)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// Creates a new user to the service
// returns an error messages correspongind to the situation
func createNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("Expected POST methcd but found: %v", r.Method)
		writeErrorResponse(w, http.StatusMethodNotAllowed, invalidMethod)
		return
	}

	var jsonUser jsons.JsonUserPassword
	reqBody, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal([]byte(reqBody), &jsonUser); err != nil {
		log.Printf("Can not parse request, %v", err)
		writeErrorResponse(w, http.StatusBadRequest, invalidRequest)
		return
	}

	err := auth.AddUser(jsonUser.Username, jsonUser.Password)
	if err != nil {
		log.Printf("Can not add user, %v", err)
		writeErrorResponse(w, http.StatusBadRequest, userAlreadyExists)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}


