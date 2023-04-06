package main

func main() {
	var myString string
	myString = "Green"
	println(myString)

	changeUsingPointer(&myString)
	println(myString)
}

func changeUsingPointer(s *string) {
	println("s is", s)
	newValue := "Red"
	println("*s is", *s)
	*s = newValue
}
