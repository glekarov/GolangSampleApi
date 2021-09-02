package main

import (
	"SampleApiServerGL/server/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a router object
	router := mux.NewRouter().StrictSlash(true)

	// Attach to the router URL controllers
	//controller.CreateServerControllers(router) // TODO: needs to be implemented

	// Attach to the router User management controllers
	controller.CreateUsersControllers(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

