package main

import "fmt"

var n int

func init() {
	fmt.Println("Init function")
	//n = 10
}

func main() {
	fmt.Println("Main function")
	fmt.Println(n)
}
