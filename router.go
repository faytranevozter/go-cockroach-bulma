package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "/static/"
)

func initRouter() *mux.Router {
	// routing
	r := mux.NewRouter()

	// static directory
	r.PathPrefix(STATIC_DIR).
		Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))))

	// first page
	r.HandleFunc("/", index_handler)

	// route user
	userrouter := r.PathPrefix("/user").Subrouter()
	userrouter.HandleFunc("/", all_user).Methods("GET")
	userrouter.HandleFunc("/get_user", get_user).Methods("GET")
	userrouter.HandleFunc("/get_detail_user/{id:[0-9]+}", get_user_detail).Methods("GET")
	userrouter.HandleFunc("/add_user", add_user).Methods("POST")
	userrouter.HandleFunc("/update_user/{id:[0-9]+}", update_user).Methods("PUT")
	userrouter.HandleFunc("/delete_user/{id:[0-9]+}", delete_user).Methods("DELETE")

	return r
}