package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Webserver is up")
}

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Server UP")
}
