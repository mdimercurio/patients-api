package main

type Store interface {
	Init() error
	GetPatient(int) *Patient
	GetPatients(int) []*Patient
	AddPatient(*Patient) error
}
