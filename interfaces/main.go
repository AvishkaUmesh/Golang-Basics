package main

import "fmt"

type Animal interface {
	Speak() string
	numLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {

	dog := Dog{
		Name:  "Fido",
		Breed: "Poodle",
	}

	gorilla := Gorilla{
		Name:  "George",
		Color: "Black",
	}

	printInfo(&dog)
	printInfo(&gorilla)

}

func printInfo(a Animal) {
	fmt.Println("This animal has", a.numLegs(), "legs and says", a.Speak())
}

func (d *Dog) Speak() string {
	return "Woof"
}

func (d *Dog) numLegs() int {
	return 4
}

func (g *Gorilla) Speak() string {
	return "Ooh Ooh Ah Ah"
}

func (g *Gorilla) numLegs() int {
	return 2
}
