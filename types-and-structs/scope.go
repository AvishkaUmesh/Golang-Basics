package main

var s = "Hello World"

func main() {
	var s2 = "Hello World 2"
	println("s is", s)
	println("s2 is", s2)

	returnTwoThings("newString")
	returnTwoThings2("newString")
}

func returnTwoThings(s string) (string, int) {
	println("s in function is", s)
	return s, 5
}

func returnTwoThings2(s3 string) (string, int) {
	println("s in function 2 is", s)
	return s3, 5
}
