package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "/static/"
	PORT       = "8888"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func all_user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is all user")
}

func user_specific(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hallo %s", vars["name"])
}

func main() {
	// routing
	r := mux.NewRouter()

	r.PathPrefix(STATIC_DIR).
        Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))))

	r.HandleFunc("/", index_handler)

	userrouter := r.PathPrefix("/user").Subrouter()
	userrouter.HandleFunc("/", all_user).Methods("GET")
	userrouter.HandleFunc("/{name}", user_specific)

	http.ListenAndServe(":" + PORT, r)
}