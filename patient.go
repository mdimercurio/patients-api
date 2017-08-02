package main

import (
	"fmt"
	"time"
)

type Patient struct {
	ID int `json:"-"`

	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthdate time.Time `json:"birthdate"`
	Sex       string    `json:"sex"`
}

func (p *Patient) Validate() error {
	if p.Email == "" {
		return fmt.Errorf("patient data missing 'email' field")
	}
	if p.FirstName == "" {
		return fmt.Errorf("patient data missing 'first_name' field")
	}
	if p.LastName == "" {
		return fmt.Errorf("patient data missing 'last_name' field")
	}
	if p.Birthdate.IsZero() {
		return fmt.Errorf("patient data missing 'birthdate' field")
	}
	if p.Sex == "" {
		return fmt.Errorf("patient data missing 'sex' field")
	}
	return nil
}
