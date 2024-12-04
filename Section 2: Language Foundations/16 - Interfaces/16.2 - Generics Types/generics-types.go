package main

import "fmt"

func genericsTypes(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	genericsTypes(1)
	genericsTypes(2)
	genericsTypes("3")
	genericsTypes(4.5)
	genericsTypes(true)
	genericsTypes(map[string]string{"a": "apple", "b": "banana"})
	genericsTypes(map[string]int{"a": 1, "b": 2})
	genericsTypes(map[string]int{"a": 1, "b": 2})

	genericsTypes(map[string]int{"a": 1, "b": 2})
}
