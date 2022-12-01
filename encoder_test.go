package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Achmad",
		MiddleName: "Sidik",
		LastName:   "Fauzi",
	}

	encoder.Encode(customer)
}
