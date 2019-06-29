package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
	"database/sql"
)

// Model
type Buku struct {
	Id string `json:"id"`
	Judul string `json:"judul"`
	Pengarang string `json:"pengarang"`
	Tahun string `json:"tahun"`
}

func getBuku() []Buku {
	rows, err := db.Query("SELECT id, judul, pengarang, tahun FROM tugas_cockroach.buku")
	if err != nil { fmt.Println(err) }
	defer rows.Close()

	x := make([]Buku, 0)
	for rows.Next() {
		var id, judul, pengarang, tahun string
		if err := rows.Scan(&id, &judul, &pengarang, &tahun); err != nil {
			fmt.Println(err)
		}
		x = append(x, Buku{id, judul, pengarang, tahun})
	}
	return x
}

// Controller
func all_buku(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/buku.html"))
	data := struct {
		Title string
	}{
		Title: "Data Buku",
	}
	tmpl.Execute(w, data)
}

func get_buku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getBuku())
	fmt.Fprintf(w, string(datajson))
}

func get_buku_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		// not valid
		fmt.Fprintf(w, "{\"success\": false, \"message\": \"ID not defined\"}")
	} else {
		var id, judul, pengarang, tahun string

		rows := db.QueryRow("SELECT id, judul, pengarang, tahun FROM tugas_cockroach.buku where id = $1", string(vars["id"]))
		err := rows.Scan(&id, &judul, &pengarang, &tahun)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Data not found\"}")
			} else {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Query Error\"}")
			}
		} else {
			datajson, _ := json.Marshal(struct {
								Success bool `json:"success"`
								Data Buku `json:"data"`
							}{Success: true, Data: Buku{id, judul, pengarang, tahun} })
			fmt.Fprintf(w, string(datajson))
		}

	}
}

func add_buku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("judul")) == 0 {
		is_success = false
	} else if len(r.FormValue("pengarang")) == 0 {
		is_success = false
	} else if len(r.FormValue("tahun")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("INSERT INTO tugas_cockroach.buku(judul, pengarang, tahun) VALUES($1, $2, $3)", 
			r.FormValue("judul"), r.FormValue("pengarang"), r.FormValue("tahun"));
		err != nil {
			fmt.Println(err)
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Buku `json:"data"`
				}{is_success, getBuku()})
	fmt.Fprintf(w, string(datajson))
}

func update_buku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	is_success := true 
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		is_success = false
	} else if len(r.FormValue("judul")) == 0 {
		is_success = false
	} else if len(r.FormValue("pengarang")) == 0 {
		is_success = false
	}else if len(r.FormValue("tahun")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("UPDATE tugas_cockroach.buku SET judul = $1, pengarang = $2, tahun = $3 WHERE id = $4", 
			r.FormValue("judul"), r.FormValue("pengarang"), r.FormValue("tahun"), string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Buku `json:"data"`
				}{is_success, getBuku()})
	fmt.Fprintf(w, string(datajson))
}

func delete_buku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	vars := mux.Vars(r)

	if len(vars["id"]) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("DELETE FROM tugas_cockroach.buku WHERE id = $1", string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
				}{is_success})
	fmt.Fprintf(w, string(datajson))
}
