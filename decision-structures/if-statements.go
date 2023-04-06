package main

func main() {
	isTrue := true

	if isTrue {
		println("isTrue is", isTrue)
	} else {
		println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		println("This is a cat")
	} else {
		println("This is not a cat")
	}

	myNumber := 10
	isTrue = false

	if myNumber > 9 && !isTrue {
		println("myNumber is greater than 9 and isTrue is", isTrue)
	} else if myNumber < 9 && !isTrue {
		println("myNumber is less than 9 and isTrue is", isTrue)
	} else {
		println("myNumber is equal to 9 and isTrue is", isTrue)
	}

}
