package main

type myStruct struct {
	Name string
}

func (m *myStruct) myFunction() {
	println("Name :", m.Name)
}

func main() {

	var myVar myStruct
	myVar.Name = "Avishka"

	myVar2 := myStruct{
		Name: "Umesh",
	}

	println("myVar :", myVar.Name)
	println("myVar2 :", myVar2.Name)

	myVar.myFunction()

}
