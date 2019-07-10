package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

const (
	STATIC_DIR = "/static/"
)

func initRouter() *mux.Router {
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
	userrouter.HandleFunc("/get_user", get_user).Methods("GET")
	userrouter.HandleFunc("/get_detail_user/{id:[0-9]+}", get_user_detail).Methods("GET")
	userrouter.HandleFunc("/add_user", add_user).Methods("POST")
	userrouter.HandleFunc("/update_user/{id:[0-9]+}", update_user).Methods("PUT")
	userrouter.HandleFunc("/delete_user/{id:[0-9]+}", delete_user).Methods("DELETE")

	// route user
	mobilruter := r.PathPrefix("/mobil").Subrouter()
	mobilruter.HandleFunc("/", all_mobil).Methods("GET")
	mobilruter.HandleFunc("/get_mobil", get_mobil).Methods("GET")
	mobilruter.HandleFunc("/get_detail_mobil/{id:[0-9]+}", get_mobil_detail).Methods("GET")
	mobilruter.HandleFunc("/add_mobil", add_mobil).Methods("POST")
	mobilruter.HandleFunc("/update_mobil/{id:[0-9]+}", update_mobil).Methods("PUT")
	mobilruter.HandleFunc("/delete_mobil/{id:[0-9]+}", delete_mobil).Methods("DELETE")

	// route buku
	bukuruter := r.PathPrefix("/buku").Subrouter()
	bukuruter.HandleFunc("/", all_buku).Methods("GET")
	bukuruter.HandleFunc("/get_buku", get_buku).Methods("GET")
	bukuruter.HandleFunc("/get_detail_buku/{id:[0-9]+}", get_buku_detail).Methods("GET")
	bukuruter.HandleFunc("/add_buku", add_buku).Methods("POST")
	bukuruter.HandleFunc("/update_buku/{id:[0-9]+}", update_buku).Methods("PUT")
	bukuruter.HandleFunc("/delete_buku/{id:[0-9]+}", delete_buku).Methods("DELETE")

	// route kucing
	kucingruter := r.PathPrefix("/kucing").Subrouter()
	kucingruter.HandleFunc("/", all_kucing).Methods("GET")
	kucingruter.HandleFunc("/get_kucing", get_kucing).Methods("GET")
	kucingruter.HandleFunc("/get_detail_kucing/{id:[0-9]+}", get_kucing_detail).Methods("GET")
	kucingruter.HandleFunc("/add_kucing", add_kucing).Methods("POST")
	kucingruter.HandleFunc("/update_kucing/{id:[0-9]+}", update_kucing).Methods("PUT")
	kucingruter.HandleFunc("/delete_kucing/{id:[0-9]+}", delete_kucing).Methods("DELETE")

	// route pegawai
	pegawairuter := r.PathPrefix("/pegawai").Subrouter()
	pegawairuter.HandleFunc("/", all_pegawai).Methods("GET")
	pegawairuter.HandleFunc("/get_pegawai", get_pegawai).Methods("GET")
	pegawairuter.HandleFunc("/get_detail_pegawai/{id:[0-9]+}", get_pegawai_detail).Methods("GET")
	pegawairuter.HandleFunc("/add_pegawai", add_pegawai).Methods("POST")
	pegawairuter.HandleFunc("/update_pegawai/{id:[0-9]+}", update_pegawai).Methods("PUT")
	pegawairuter.HandleFunc("/delete_pegawai/{id:[0-9]+}", delete_pegawai).Methods("DELETE")

	return r
}