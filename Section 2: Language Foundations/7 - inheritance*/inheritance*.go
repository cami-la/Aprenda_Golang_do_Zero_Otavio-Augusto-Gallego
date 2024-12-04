package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	person
	course string
}

func main() {
	fmt.Println("inheritance")

	person1 := person{"João", "Pedro", 20, 178}
	fmt.Println(person1)

	student := student{person1, "Engenharia"}

	fmt.Println(student.height)
}
