package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Achmad",
		MiddleName: "Sidik",
		LastName:   "Fauzi",
		Age:        23,
		Married:    false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
