package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Aplle Macbook Air M2","imageUrl":"http://example.com/image.png", "price": 23600000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["imageUrl"])
	fmt.Println(result["price"])
}

func TestMap(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apple Macbook Air M2",
		"price": 23600000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
