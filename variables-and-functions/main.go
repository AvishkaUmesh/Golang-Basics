package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var num int

	whatToSay = "Hello, World!"
	fmt.Println(whatToSay)

	num = 10
	fmt.Println("number is set to", num)

	someThing := saySomeThing()
	fmt.Println("The function returned", someThing)

	someThing2, someOtherThing := returnTwoThings()
	fmt.Println("The function returned", someThing2, "and", someOtherThing)

	name, age := "John", 30
	fmt.Println("My name is", name, "and I am", age, "years old")

}

func saySomeThing() string {
	return "something"
}

func returnTwoThings() (string, int) {
	return "something", 5
}
