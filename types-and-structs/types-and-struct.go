package main

import "time"

type User struct {
	Name     string
	Age      int
	Phone    int
	Birthday time.Time
}

// if a variable name is capitalized, it is exported and can be used outside the package

func main() {
	user := User{
		Name:  "Avishka",
		Age:   21,
		Phone: 123456789,
	}

	println("Name :", user.Name)
	println("Age :", user.Age)
	println("Phone :", user.Phone)
}
