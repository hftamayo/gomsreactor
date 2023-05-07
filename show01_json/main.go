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
	decodeIntoMap()
	encodeSlideOfStructs()
	encodingJsonWithIndent()
	decodeIntoSlice()
	encodeStructWithOmitEmpty()
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
	fmt.Println("\nDecode JSON:")
	fmt.Println(person)
}

func decodeIntoMap(){
	jsonData := []byte(`{"name": "Liam", "age": 55}`)

	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDecode JSON into map:")
	fmt.Println(data)
}

func encodeSlideOfStructs(){
	person1 := Person{Name: "John", Age: 32}
	person2 := Person{Name: "Liam", Age: 55}
	people := []Person{person1, person2}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncode slice of structs:")
	fmt.Println(string(jsonData))
}

func encodingJsonWithIndent(){
	person := Person{Name: "John", Age: 32}

	jsonData, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncoding JSON with indent: ")
	fmt.Println(string(jsonData))

}

func decodeIntoSlice(){
	jsonData := []byte(`[{"name": "John", "age": 32}, {"name": "Liam", "age": 55}]`)

	var people []Person
	err := json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nDecode JSON into Slice: ")
	for _, person := range people {
		fmt.Printf("Name: %s\nAge: %d\n", person.Name, person.Age)
	}
}

func encodeStructWithOmitEmpty(){
	type Person struct {
		Name string `json:"name"`
		Age int 	`json: "age,omitempty`
	}

	person := Person{Name: "John"}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEncode struct with omitempty: ")
	fmt.Println(string(jsonData))
}