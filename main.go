package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PatientsPerPage = 10

func main() {
	log.Println("Starting server")

	store := NewMemoryStore()
	if err := store.Init(); err != nil {
		log.Fatalln(err)
	}
	server := NewServer(store)

	router := mux.NewRouter()
	router.Path("/v1/patients/{id}").Methods("GET").HandlerFunc(server.GetPatient)
	router.Path("/v1/patients").Methods("GET").HandlerFunc(server.GetPatients)
	router.Path("/v1/patients").Methods("POST").HandlerFunc(server.PostPatient)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
}
