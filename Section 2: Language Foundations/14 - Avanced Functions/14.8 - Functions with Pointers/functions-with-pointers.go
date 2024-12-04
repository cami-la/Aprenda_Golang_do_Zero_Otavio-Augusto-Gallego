package main

import "fmt"

func invertSignal(n int) int {
	return n * -1
}

func invertSignalPointer(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	fmt.Println(n)
	invertSignalPointer(&n)
	fmt.Println(n)

	n = 20
	fmt.Println(n)
	newNumber := invertSignal(n)
	fmt.Println(n)
	fmt.Println(newNumber)
}
