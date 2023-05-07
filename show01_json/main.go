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