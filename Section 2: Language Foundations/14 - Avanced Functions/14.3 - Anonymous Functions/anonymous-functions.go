package main

import "fmt"

func main() {
	resullt := func(text string) string {
		return fmt.Sprintf("Returned %s", text)
	}("Hello")

	fmt.Println(resullt)
}
