package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CreateServerControllers(router *mux.Router) {
	router.HandleFunc("/urls", storeUrlList).Methods("POST")
	router.HandleFunc("/urls", returnAllUrls)
}

func returnAllUrls(w http.ResponseWriter, r *http.Request) {
	// TODO: will be implemented later
}

func storeUrlList(w http.ResponseWriter, r *http.Request) {
	// TODO: will be implemented later
}

