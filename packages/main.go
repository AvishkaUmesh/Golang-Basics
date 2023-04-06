package main

import (
	"fmt"

	"github.com/AvishkaUmesh/Golang-Basics/packages/helpers"
)

func main() {

	var myVar helpers.SomeType
	myVar.Name = "Avishka"
	myVar.Number = 123

	fmt.Println("Name: ", myVar.Name)
	fmt.Println("Number: ", myVar.Number)
}
