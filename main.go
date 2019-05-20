package main

import (
	"fmt"
	"os"
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
	if len(os.Getenv("ROACH_USER")) == 0 { os.Setenv("ROACH_USER", "root") }
	if len(os.Getenv("ROACH_HOST")) == 0 { os.Setenv("ROACH_HOST", "localhost") }
	if len(os.Getenv("ROACH_PORT")) == 0 { os.Setenv("ROACH_PORT", "26257") }
	if len(os.Getenv("ROACH_DB")) == 0 { os.Setenv("ROACH_DB", "tugas_cockroach") }

	db, err = sql.Open("postgres",
		"postgresql://" + os.Getenv("ROACH_USER") + "@" + os.Getenv("ROACH_HOST") + ":" + os.Getenv("ROACH_PORT") + "/" + os.Getenv("ROACH_DB") + "?sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to the database: ", err)
	}
	defer db.Close()

	// init router mux (router.go)
	r := initRouter()

	// serve at port x
	http.ListenAndServe(":" + PORT, r)
}