package main

import "fmt"

func main() {
	fmt.Println("Pointers Go")

	var variable int = 10
	var variable2 int = variable
	fmt.Println(variable, variable2)

	variable2++
	fmt.Println(variable, variable2)

	//Pointer is a reference to a memory address
	var variable3 int
	var pointer *int
	fmt.Println(variable3, pointer)

	variable3 = 100
	pointer = &variable3
	fmt.Println(variable3, pointer) // dereference

	variable3 = 150
	fmt.Println(variable3, *pointer)
}
