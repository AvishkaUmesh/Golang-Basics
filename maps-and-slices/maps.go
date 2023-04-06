package main

type Person struct {
	Name string
	Age  int
}

func main() {

	myMap := make(map[string]string)

	myMap["name"] = "Avishka"
	myMap["age"] = "25"

	println(myMap["name"])
	println(myMap["age"])

	println("----------------------------")

	myMap2 := make(map[string]int)

	myMap2["age"] = 25
	myMap2["count"] = 10

	println(myMap2["age"])
	println(myMap2["count"])

	println("----------------------------")

	myMap3 := make(map[string]Person)

	myMap3["person1"] = Person{"Avishka", 25}
	myMap3["person2"] = Person{"Umesh", 30}

	println(myMap3["person1"].Name)
	println(myMap3["person1"].Age)
	println("----------------------------")
	println(myMap3["person2"].Name)
	println(myMap3["person2"].Age)

}
