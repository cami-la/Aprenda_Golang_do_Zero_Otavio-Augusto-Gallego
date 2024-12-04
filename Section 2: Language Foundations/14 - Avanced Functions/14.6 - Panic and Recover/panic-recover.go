package main

import "fmt"

func recoverFunction() {
	if r := recover(); r != nil {
		fmt.Println("Recovered")
	}
}

func aprovedOrNot(grade float64) string {
	defer recoverFunction()
	if grade >= 6 {
		return "Approved"
	}
	panic("Disapproved")
}

func main() {
	fmt.Println(aprovedOrNot(7))
	fmt.Println(aprovedOrNot(5))
}
