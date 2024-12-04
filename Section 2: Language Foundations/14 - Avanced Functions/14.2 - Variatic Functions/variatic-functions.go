package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, values ...int) {
	for _, value := range values {
		fmt.Println(text, value)
	}
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)

	fmt.Println("-----------------")
	write("Number", 1, 2, 3, 4, 5)
}
