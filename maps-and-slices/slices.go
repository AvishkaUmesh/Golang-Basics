package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var myStringSlice []string

	myStringSlice = append(myStringSlice, "Avishka")
	myStringSlice = append(myStringSlice, "Umesh")

	println(myStringSlice[0])
	println(myStringSlice[1])

	println(myStringSlice) // prints the memory address of the slice

	// fmt.Println() and log.Println() will print the contents of the slice
	fmt.Println(myStringSlice)
	log.Println(myStringSlice)

	var myIntSlice []int

	myIntSlice = append(myIntSlice, 30)
	myIntSlice = append(myIntSlice, 10)
	myIntSlice = append(myIntSlice, 20)

	sort.Ints(myIntSlice)

	fmt.Println(myIntSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // index 0 to 9

	fmt.Println(numbers)
	fmt.Println(numbers[0:2])  // prints the first two elements
	fmt.Println(numbers[6:10]) // prints the last four elements
	fmt.Println(numbers[9])

}
