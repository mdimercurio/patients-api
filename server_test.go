package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	store := NewMemoryStore()
	if err := store.Init(); err != nil {
		log.Fatalln(err)
	}
	server := NewServer(store)

	router := mux.NewRouter()
	router.Path("/v1/patients").Methods("POST").HandlerFunc(server.PostPatient)

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatalln(err)
		}
	}()
}

func BenchmarkPost(b *testing.B) {
	for n := 0; n < b.N; n++ {
		body := bytes.NewBufferString(fmt.Sprintf(`{"email":"%s","first_name":"bob","last_name":"bobby","birthdate":"2000-01-01T00:00:00Z","sex":"Male"}`, uuid.New().String()))
		resp, err := http.Post("http://localhost:8080/v1/patients", "application/json", body)
		require.NoError(b, err)
		assert.Equal(b, 201, resp.StatusCode)
	}
}
