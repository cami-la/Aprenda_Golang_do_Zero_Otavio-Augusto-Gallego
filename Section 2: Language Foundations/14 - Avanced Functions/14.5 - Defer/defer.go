package main

import "fmt"

func function1() {
	fmt.Println("First")
}

func function2() {
	fmt.Println("Second")
}

func aprovedOrNot(grade float64) string {
	defer fmt.Println("End")
	if grade >= 6 {
		return "Approved"
	}
	return "Disapproved"
}

func main() {
	defer function1()
	function2()

	fmt.Println(aprovedOrNot(7))
	fmt.Println(aprovedOrNot(5))
}
