package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck).Methods("GET")
	r.HandleFunc("/db", DbCheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Webserver is up")
	log.Println("API server is up")
}

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Server UP")
	log.Println("HealthChecked")
}

func DbCheck(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "postgres://postgres:password@172.17.0.6/alerts?sslmode=disable")
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 400)
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 400)
	}
	fmt.Fprint(w, "Can ping db")
}
