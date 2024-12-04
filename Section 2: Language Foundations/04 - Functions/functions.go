package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculates(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sumResult := sum(10, 20)
	fmt.Println(sumResult)

	var f = func(txt string) string {
		 fmt.Println("Func f")
		 return txt
	}

	result := f("Camila")

	println(result)

	resultSum, resultSub := calculates(10, 15)
	//resultSum, _ := calculates(10, 15)
	fmt.Println(resultSum, resultSub)
}
