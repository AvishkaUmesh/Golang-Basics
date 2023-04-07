package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `[{
	"first_name": "Clark",
	"last_name": "Kent",
	"hair_color": "black",
	"has_dog": true
	}]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println(err)
	}

	for _, person := range unmarshalled {
		println(person.FirstName)
		println(person.LastName)
		println(person.HairColor)
		println(person.HasDog)
	}
}
