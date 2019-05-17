package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	PORT = "8888"
)

var db *sql.DB
var err error

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
	// connect to the database
	db, err = sql.Open("postgres",
		"postgresql://root@localhost:26257/tugas_cockroach?sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to the database: ", err)
	}
	defer db.Close()

	// init router mux (router.go)
	r := initRouter()

	// serve at port x
	http.ListenAndServe(":" + PORT, r)
}