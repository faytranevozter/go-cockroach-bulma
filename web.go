package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

const (
	STATIC_DIR = "/static/"
	PORT       = "8888"
)

type DataUser struct {
	Number int
	Id string
	Name string
	Email string
}

type PageUsers struct {
	Title string
	Data [] DataUser 
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/users.html"))
	data := PageUsers {
		Title: "Halaman Home",
		Data: [] DataUser {
			{Number: 1, Id: "1", Name: "Fahrur", Email: "fahrur@mail.com"},
			{Number: 2, Id: "3", Name: "Laily", Email: "laily@mail.com"},
			{Number: 3, Id: "5", Name: "Oky", Email: "oky@mail.com"},
			{Number: 4, Id: "6", Name: "Indra", Email: "indra@mail.com"},
			{Number: 5, Id: "7", Name: "Rismi", Email: "rismi@mail.com"},
		},
	}
	tmpl.Execute(w, data)
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

	// static directory
	r.PathPrefix(STATIC_DIR).
        Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("." + STATIC_DIR))))

    // first page
	r.HandleFunc("/", index_handler)

	// testing sub router
	userrouter := r.PathPrefix("/user").Subrouter()
	userrouter.HandleFunc("/", all_user).Methods("GET")
	userrouter.HandleFunc("/{name}", user_specific)

	http.ListenAndServe(":" + PORT, r)
}