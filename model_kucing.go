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
type Kucing struct {
	Id string `json:"id"`
	Ras string `json:"ras"`
	Negara string `json:"negara"`
	Warna string `json:"warna"`
}

func getKucing() []Kucing {
	rows, err := db.Query("SELECT id, ras, negara, warna FROM tugas_cockroach.kucing")
	if err != nil { fmt.Println(err) }
	defer rows.Close()

	x := make([]Kucing, 0)
	for rows.Next() {
		var id, ras, negara, warna string
		if err := rows.Scan(&id, &ras, &negara, &warna); err != nil {
			fmt.Println(err)
		}
		x = append(x, Kucing{id, ras, negara, warna})
	}
	return x
}

// Controller
func all_kucing(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/kucing.html"))
	data := struct {
		Title string
	}{
		Title: "Data Kucing",
	}
	tmpl.Execute(w, data)
}

func get_kucing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getKucing())
	fmt.Fprintf(w, string(datajson))
}

func get_kucing_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		// not valid
		fmt.Fprintf(w, "{\"success\": false, \"message\": \"ID not defined\"}")
	} else {
		var id, ras, negara, warna string

		rows := db.QueryRow("SELECT id, ras, negara, warna FROM tugas_cockroach.kucing where id = $1", string(vars["id"]))
		err := rows.Scan(&id, &ras, &negara, &warna)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Data not found\"}")
			} else {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Query Error\"}")
			}
		} else {
			datajson, _ := json.Marshal(struct {
								Success bool `json:"success"`
								Data Kucing `json:"data"`
							}{Success: true, Data: Kucing{id, ras, negara, warna} })
			fmt.Fprintf(w, string(datajson))
		}

	}
}

func add_kucing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("ras")) == 0 {
		is_success = false
	} else if len(r.FormValue("negara")) == 0 {
		is_success = false
	}else if len(r.FormValue("warna")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("INSERT INTO tugas_cockroach.kucing(ras, negara, warna) VALUES($1, $2, $3)", 
			r.FormValue("ras"), r.FormValue("negara"), r.FormValue("warna"));
		err != nil {
			fmt.Println(err)
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Kucing `json:"data"`
				}{is_success, getKucing()})
	fmt.Fprintf(w, string(datajson))
}

func update_kucing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	is_success := true 
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		is_success = false
	} else if len(r.FormValue("ras")) == 0 {
		is_success = false
	} else if len(r.FormValue("negara")) == 0 {
		is_success = false
	} else if len(r.FormValue("warna")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("UPDATE tugas_cockroach.kucing SET ras = $1, negara = $2, warna = $3 WHERE id = $4", 
			r.FormValue("ras"), r.FormValue("negara"), r.FormValue("warna"), string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Kucing `json:"data"`
				}{is_success, getKucing()})
	fmt.Fprintf(w, string(datajson))
}

func delete_kucing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	vars := mux.Vars(r)

	if len(vars["id"]) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("DELETE FROM tugas_cockroach.kucing WHERE id = $1", string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
				}{is_success})
	fmt.Fprintf(w, string(datajson))
}
