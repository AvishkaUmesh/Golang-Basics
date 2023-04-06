package main

func main() {

	myVar := "dog"

	switch myVar {
	case "cat":
		println("This is a cat")
	case "dog":
		println("This is a dog")
	default:
		println("This is not a cat or a dog")
	}

}
