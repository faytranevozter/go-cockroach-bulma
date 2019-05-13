package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
)

const (
	STATIC_DIR = "/static/"
	PORT       = "8888"
)

/* USER */
type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func getUser() []User {
	return [] User {
		{Id: "1", Name: "Fahrur", Email: "fahrur@mail.com"},
		{Id: "3", Name: "Laily", Email: "laily@mail.com"},
		{Id: "5", Name: "Oky", Email: "oky@mail.com"},
		{Id: "6", Name: "Indra", Email: "indra@mail.com"},
		{Id: "7", Name: "Rismi", Email: "rismi@mail.com"},
	}
}

func all_user(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/user.html"))
	data := struct {
		Title string
	}{
		Title: "Data User",
	}
	tmpl.Execute(w, data)
}

func get_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getUser())
	fmt.Fprintf(w, string(datajson))
}

func add_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("name")) == 0 {
		is_success = false
	} else if len(r.FormValue("email")) == 0 {
		is_success = false
	}
	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []User `json:"data"`
				}{is_success, getUser()})
	fmt.Fprintf(w, string(datajson))
}

func update_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getUser())
	fmt.Fprintf(w, string(datajson))
}

func delete_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getUser())
	fmt.Fprintf(w, string(datajson))
}
/* END OF USER */

// main handler
func index_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	data := struct {
		Title string
	}{
		Title: "Halaman Home",
	}
	tmpl.Execute(w, data)
}

func main() {
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
	userrouter.HandleFunc("/get", get_user).Methods("GET")
	userrouter.HandleFunc("/add", add_user).Methods("POST")
	userrouter.HandleFunc("/update", update_user).Methods("PUT")
	userrouter.HandleFunc("/delete", delete_user).Methods("DELETE")

	http.ListenAndServe(":" + PORT, r)
}