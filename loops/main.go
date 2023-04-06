package main

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {

	for i := 0; i <= 5; i++ {
		println(i)
	}

	println("--------------------------")

	animals := []string{"dog", "cat", "bird"}

	for i := 0; i < len(animals); i++ {
		println(animals[i])
	}

	println("--------------------------")
	// range over
	for i, animal := range animals {
		println(i, animal)
	}

	println("--------------------------")

	for _, animal := range animals {
		println(animal)
	}

	animalsMap := make(map[string]string)
	animalsMap["dog"] = "woof"
	animalsMap["cat"] = "meow"
	animalsMap["bird"] = "tweet"

	for animal, animalSound := range animalsMap {
		println(animal, animalSound)
	}

	println("--------------------------")

	var users []User
	users = append(users, User{"John", 20, "test@test.com"})
	users = append(users, User{"Jane", 30, "test@test.com"})

	for _, user := range users {
		println(user.Name, user.Age, user.Email)
	}

}
