package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
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
	rows, err := db.Query("SELECT id, name, email FROM tugas_cockroach.user")
	if err != nil { fmt.Println(err) }
	defer rows.Close()

	x := make([]User, 0)
	for rows.Next() {
		var id, name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			fmt.Println(err)
		}
		x = append(x, User{id, name, email})
	}
	return x
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

func get_user_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if len(r.FormValue("id")) == 0 {
		// not valid
		fmt.Fprintf(w, "{\"success\": false, \"message\": \"ID not defined\"}")
	} else {
		var id, name, email string

		rows := db.QueryRow("SELECT id, name, email FROM tugas_cockroach.user where id = $1", string(r.FormValue("id")))
		err := rows.Scan(&id, &name, &email)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Data not found\"}")
			} else {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Query Error\"}")
			}
		} else {
			datajson, _ := json.Marshal(struct {
								Success bool `json:"success"`
								Data User `json:"data"`
							}{Success: true, Data: User{id, name, email} })
			fmt.Fprintf(w, string(datajson))
		}

	}
}

func add_user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("name")) == 0 {
		is_success = false
	} else if len(r.FormValue("email")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("INSERT INTO tugas_cockroach.user(name, email) VALUES('$1', '$2')", 
			r.FormValue("name"), r.FormValue("email"));
		err != nil {
			is_success = false
		}
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

var db *sql.DB
var err error

func main() {
	// connect to the database
	db, err = sql.Open("postgres",
		"postgresql://root@localhost:26257/tugas_cockroach?sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to the database: ", err)
	}
	defer db.Close()

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
	userrouter.HandleFunc("/get_detail", get_user_detail).Methods("GET")
	userrouter.HandleFunc("/add", add_user).Methods("POST")
	userrouter.HandleFunc("/update", update_user).Methods("PUT")
	userrouter.HandleFunc("/delete", delete_user).Methods("DELETE")

	http.ListenAndServe(":" + PORT, r)
}