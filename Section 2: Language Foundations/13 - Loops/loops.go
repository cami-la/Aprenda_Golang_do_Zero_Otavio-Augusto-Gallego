package main

import (
	"fmt"
)

func main() {
	//i := 0

	/*for i < 10 {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
	}*/

	//fmt.Println(i)

	/*for j := 0; j < 10; j += 2 {
		fmt.Println("Increment J", j)
		time.Sleep(time.Second)
	}*/

	names := []string{"John", "Doe", "Jane", "Doe"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	for _, value := range "WORD" {
		fmt.Println(string(value))
	}

	for key, value := range map[string]string{"a": "apple", "b": "banana"} {
		fmt.Println(key, value)
	}

	// Infinite loop
	/*for {
	fmt.Println("Infinite loop")
	*/
}
