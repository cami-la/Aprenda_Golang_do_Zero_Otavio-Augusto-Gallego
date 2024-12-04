package main

import "fmt"

func closure() func() {
	text := "Inside closure"
	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	function := closure()
	function()
}
