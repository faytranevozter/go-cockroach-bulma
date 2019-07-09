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
type Mobil struct {
	Id string `json:"id"`
	Mobil string `json:"mobil"`
	Jenis string `json:"jenis"`
	Tahun string `json:"tahun"`
}

func getMobil() []Mobil {
	rows, err := db.Query("SELECT id, mobil, jenis, tahun FROM tugas_cockroach.mobil")
	if err != nil { fmt.Println(err) }
	defer rows.Close()

	x := make([]Mobil, 0)
	for rows.Next() {
		var id, mobil, jenis, tahun string
		if err := rows.Scan(&id, &mobil, &jenis, &tahun); err != nil {
			fmt.Println(err)
		}
		x = append(x, Mobil{id, mobil, jenis, tahun})
	}

	
	return x
}

// Controller
func all_mobil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/mobil.html"))
	data := struct {
		Title string
	}{
		Title: "Data Mobil",
	}
	tmpl.Execute(w, data)
}

func get_mobil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	datajson, _ := json.Marshal(getMobil())
	fmt.Fprintf(w, string(datajson))
}

func get_mobil_detail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	if len(vars["id"]) == 0 {
		// not valid
		fmt.Fprintf(w, "{\"success\": false, \"message\": \"ID not defined\"}")
	} else {
		var id, mobil, jenis, tahun string

		rows := db.QueryRow("SELECT id, mobil, jenis, tahun FROM tugas_cockroach.mobil where id = $1", string(vars["id"]))
		err := rows.Scan(&id, &mobil, &jenis, &tahun)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Data not found\"}")
			} else {
				fmt.Fprintf(w, "{\"success\": false, \"message\": \"Query Error\"}")
			}
		} else {
			datajson, _ := json.Marshal(struct {
				Success bool `json:"success"`
				Data Mobil `json:"data"`
				}{Success: true, Data: Mobil{id, mobil, jenis, tahun} })
			fmt.Fprintf(w, string(datajson))
		}

	}
}

func add_mobil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	is_success := true 
	if len(r.FormValue("mobil")) == 0 {
		is_success = false
		} else if len(r.FormValue("jenis")) == 0 {
			is_success = false
			} else if len(r.FormValue("tahun")) == 0 {
				is_success = false
			}

			if is_success {
				if _, err := db.Exec("INSERT INTO tugas_cockroach.mobil(mobil, jenis, tahun) VALUES($1, $2, $3)", 
					r.FormValue("mobil"), r.FormValue("jenis"), r.FormValue("tahun"));
				err != nil {
					fmt.Println(err)
					is_success = false
				}
			}

			datajson, _ := json.Marshal(struct {
				Success bool `json:"success"`
				Data []Mobil `json:"data"`
				}{is_success, getMobil()})
			fmt.Fprintf(w, string(datajson))
		}

		func update_mobil(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")

			is_success := true 
			vars := mux.Vars(r)
			if len(vars["id"]) == 0 {
				is_success = false
				} else if len(r.FormValue("mobil")) == 0 {
					is_success = false
					} else if len(r.FormValue("jenis")) == 0 {
						is_success = false
					} else if len(r.FormValue("tahun")) == 0 {
						is_success = false
					}

					if is_success {
						if _, err := db.Exec("UPDATE tugas_cockroach.mobil SET mobil = $1, jenis = $2, tahun = $3 WHERE id = $4", 
							r.FormValue("mobil"), r.FormValue("jenis"), r.FormValue("tahun"), string(vars["id"]));
						err != nil {
							is_success = false
						}
					}

					datajson, _ := json.Marshal(struct {
						Success bool `json:"success"`
						Data []Mobil `json:"data"`
						}{is_success, getMobil()})
					fmt.Fprintf(w, string(datajson))
				}

				func delete_mobil(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-type", "application/json")
					is_success := true 
					vars := mux.Vars(r)

					if len(vars["id"]) == 0 {
						is_success = false
					}

					if is_success {
						if _, err := db.Exec("DELETE FROM tugas_cockroach.mobil WHERE id = $1", string(vars["id"]));
						err != nil {
							is_success = false
						}
					}

					datajson, _ := json.Marshal(struct {
						Success bool `json:"success"`
						}{is_success})
					fmt.Fprintf(w, string(datajson))
				}
