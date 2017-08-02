package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

type MemoryStore struct {
	filePath string
	mutex    sync.Mutex

	patients []*Patient
	file     *os.File
}

func NewMemoryStore() *MemoryStore {
	path := os.Getenv("DATA_PATH")
	if path == "" {
		path = "patients.json"
	}
	return &MemoryStore{
		filePath: path,
	}
}

func (s *MemoryStore) Init() error {
	f, err := os.OpenFile(s.filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	patients := make([]*Patient, 0)
	for scanner.Scan() {
		var p *Patient
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		if err = json.Unmarshal(b, &p); err != nil {
			return err
		}

		patients = append(patients, p)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	s.patients = patients
	return nil
}

func (s *MemoryStore) GetPatient(id int) *Patient {
	if id >= len(s.patients) {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.patients[id]
}

func (s *MemoryStore) GetPatients(page int) []*Patient {
	out := make([]*Patient, 0)
	firstID := page * PatientsPerPage

	for id := firstID; id < firstID+PatientsPerPage; id++ {
		p := s.GetPatient(id)
		if p == nil {
			break
		}

		out = append(out, p)
	}

	return out
}

var ErrAlreadyExists = errors.New("patient already exists")

func (s *MemoryStore) AddPatient(in *Patient) error {
	for _, p := range s.patients {
		if p.Email == in.Email {
			return ErrAlreadyExists
		}
	}

	s.mutex.Lock()
	s.patients = append(s.patients, in)
	s.mutex.Unlock()

	s.writePatient(in)
	return nil
}

func (s *MemoryStore) writePatient(p *Patient) {
	f, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	if err := enc.Encode(p); err != nil {
		log.Fatalln(err)
	}
}
