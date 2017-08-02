package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	Store Store
}

func NewServer(s Store) *Server {
	return &Server{
		Store: s,
	}
}

type PatientsResp struct {
	Data  []*Patient `json:"data"`
	Links Pagination `json:"links"`
}

func (s *Server) GetPatients(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("page"))
	patients := s.Store.GetPatients(page)

	data := PatientsResp{
		Data:  patients,
		Links: newPagination(page, len(patients)),
	}
	respondJSON(w, 200, data)
}

func (s *Server) GetPatient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		httpNotFound(w)
		return
	}

	patient := s.Store.GetPatient(id)
	if patient == nil {
		httpNotFound(w)
		return
	}

	respondJSON(w, 200, patient)
}

func (s *Server) PostPatient(w http.ResponseWriter, r *http.Request) {
	patient := &Patient{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&patient); err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	if err := patient.Validate(); err != nil {
		httpError(w, err, 400, "invalid_patient_data")
		return
	}

	err := s.Store.AddPatient(patient)
	if err != nil {
		if err == ErrAlreadyExists {
			httpError(w, err, 409, "patient_already_exists")
			return
		}
		log.Println(err)
		return
	}

	respondJSON(w, 201, patient)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		log.Println(err)
	}
}

func httpNotFound(w http.ResponseWriter) {
	httpError(w, nil, 404, "not_found")
}

type httpErr struct {
	ID     string
	Status string
	Title  string
	Detail string
	Code   string
	Source map[string]interface{}
}

func httpError(w http.ResponseWriter, err error, status int, code string) {
	id := uuid.New().String()
	statusStr := fmt.Sprintf("%d", status)
	var data = struct {
		Errors []httpErr
	}{
		Errors: []httpErr{
			httpErr{
				ID:     id,
				Status: statusStr,
				Title:  http.StatusText(status),
				Code:   code,
				Detail: err.Error(),
			},
		},
	}

	respondJSON(w, status, data)

	log.Printf("%s: %s %s %q\n", id, statusStr, code)
}
