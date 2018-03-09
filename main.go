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
	log.Infof("API server is up")
}

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Server UP")
	log.Info("HealthChecked")
}

func DbCheck(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/alerts?sslmode=disable")
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 400)
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 400)
	}
	fmt.Fprintln(w, "Can ping db")

}
