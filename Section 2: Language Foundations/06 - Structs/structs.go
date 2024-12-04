package main

import "fmt"

type user struct {
	name string
	age  uint8
	adress adress
}

type adress struct {
	street string
	number uint8
}

func main() {
	fmt.Println("File Struct")

	fmt.Println("----------------------")
	var user1 user
	user1.name = "cami"
	user1.age = 30
	fmt.Println(user1)

	adressEx := adress{street: "street1", number: 70}

	fmt.Println("----------------------")
	user2 := user{"cami2", 30, adressEx}
	fmt.Println(user2)

	fmt.Println("----------------------")
	user3 := user{name: "Cami"}
	fmt.Println(user3)

	fmt.Println("----------------------")
	user4 := user{age: 30}
	fmt.Println(user4)
}
