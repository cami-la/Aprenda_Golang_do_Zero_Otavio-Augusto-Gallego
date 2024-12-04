package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10
	if number > 15 {
		fmt.Println("Number is bigger than 15")
	} else {
		fmt.Println("Number is smaller than 15")
	}

	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Number is bigger than 15")
	} else if anotherNumber < 0 {
		fmt.Println("Number is smaller than 15")
	} else {
		fmt.Println("Number is 0")
	}
}
