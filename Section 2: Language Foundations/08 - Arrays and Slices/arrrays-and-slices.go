package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string
	array1[0] = "Pos 1"
	fmt.Println(array1)

	array2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Changed"
	fmt.Println(array2)
	fmt.Println(slice2)

	//Intern Array
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println("-----------------")
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
