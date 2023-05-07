package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	encodeJson()
	decodeJson()
}

func encodeJson(){
	person := Person{Name: "John", Age: 32}

	jsonData, err:= json.Marshal(person)
	if err!= nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Encode JSON:")
	fmt.Println(string(jsonData))
}

func decodeJson(){
	jsonData := []byte(`{"name": "John", "age": 32}`)

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
}