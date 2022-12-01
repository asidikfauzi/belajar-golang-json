package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Achmad",
		MiddleName: "Sidik",
		LastName:   "Fauzi",
		Age:        23,
		Married:    false,
		Hobbies:    []string{"Futsal", "Traveling", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Achmad","MiddleName":"Sidik","LastName":"Fauzi","Age":23,"Married":false,"Hobbies":["Futsal","Traveling","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}
