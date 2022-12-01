package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Aplle Macbook Air M2",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Aplle Macbook Air M2","imageUrl":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
