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
type Pegawai struct {
	Id string `json:"id"`
	Kode string `json:"kdpeg"`
	Nama string `json:"nama"`
	Jabatan string `json:"jabatan"`
	Divisi string `json:"divisi"`
}

func getPegawai() []Pegawai {
	rows, err := db.Query("SELECT id, kdpeg, nama, jabatan, divisi FROM tugas_cockroach.pegawai")
	if err != nil { fmt.Println(err) }
	defer rows.Close()

	x := make([]Pegawai, 0)
	for rows.Next() {
		var id, kdpeg, nama, jabatan, divisi string
		if err := rows.Scan(&id, &kdpeg, &nama, &jabatan, &divisi); err != nil {
			fmt.Println(err)
		}
		x = append(x, Pegawai{id, kdpeg, nama, jabatan, divisi})
	}
	return x
}

// Controller
func all_pegawai(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/pegawai.html"))
	data := struct {
		Title string
	}{
		Title: "Data Pegawai",
	}
	tmpl.Execute(w, data)
}

func get_pegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getPegawai())
	fmt.Fprintf(w, string(datajson))
}

func get_pegawai_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		// not valid
		fmt.Fprintf(w, "{\"success\": false, \"message\": \"ID not defined\"}")
	} else {
		var id, kdpeg, nama, jabatan, divisi string

		rows := db.QueryRow("SELECT id, kdpeg, nama, jabatan, divisi FROM tugas_cockroach.pegawai where id = $1", string(vars["id"]))
		err := rows.Scan(&id, &kdpeg, &nama, &jabatan, &divisi)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Data tidak ditemukan\"}")
			} else {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Query Error\"}")
			}
		} else {
			datajson, _ := json.Marshal(struct {
								Success bool `json:"success"`
								Data Pegawai `json:"data"`
							}{Success: true, Data: Pegawai{id, kdpeg, nama, jabatan, divisi} })
			fmt.Fprintf(w, string(datajson))
		}

	}
}

func add_pegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("kdpeg")) == 0 {
		is_success = false
	} else if len(r.FormValue("nama")) == 0 {
		is_success = false
	} else if len(r.FormValue("jabatan")) == 0 {
		is_success = false
	} else if len(r.FormValue("divisi")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("INSERT INTO tugas_cockroach.pegawai(kdpeg, nama, jabatan, divisi) VALUES($1, $2, $3, $4)", 
			r.FormValue("kdpeg"), r.FormValue("nama"), r.FormValue("jabatan"), r.FormValue("divisi"));
		err != nil {
			fmt.Println(err)
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Pegawai `json:"data"`
				}{is_success, getPegawai()})
	fmt.Fprintf(w, string(datajson))
}

func update_pegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	is_success := true 
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		is_success = false
	} else if len(r.FormValue("kdpeg")) == 0 {
		is_success = false
	} else if len(r.FormValue("nama")) == 0 {
		is_success = false
	} else if len(r.FormValue("jabatan")) == 0 {
		is_success = false
	} else if len(r.FormValue("divisi")) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("UPDATE tugas_cockroach.pegawai SET kdpeg = $1, nama = $2, jabatan = $3, divisi = $4 WHERE id = $5", 
			r.FormValue("kdpeg"), r.FormValue("nama"), r.FormValue("jabatan"), r.FormValue("divisi"), string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
					Data []Pegawai `json:"data"`
				}{is_success, getPegawai()})
	fmt.Fprintf(w, string(datajson))
}

func delete_pegawai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	vars := mux.Vars(r)

	if len(vars["id"]) == 0 {
		is_success = false
	}

	if is_success {
		if _, err := db.Exec("DELETE FROM tugas_cockroach.pegawai WHERE id = $1", string(vars["id"]));
		err != nil {
			is_success = false
		}
	}

	datajson, _ := json.Marshal(struct {
					Success bool `json:"success"`
				}{is_success})
	fmt.Fprintf(w, string(datajson))
}
