package main

import "fmt"

func main() {
	//ARITIMETICS
	fmt.Println("----------------------")

	sum := 1 + 2
	sub := 1 - 2
	division := 10 / 4
	mult := 10 * 5
	mod := 10 % 2
	fmt.Println(sum, sub, division, mult, mod)

	fmt.Println("----------------------")
	var number1 int16 = 10
	var number2 int16 = 25
	sum2 := number1 + number2
	fmt.Println(sum2)

	//END ARITHIMETICS

	//ATRIBUTES
	fmt.Println("----------------------")
	var var1 string = "String"
	var2 := "String2"
	fmt.Println(var1, var2)

	//END ATRIBUTES

	//RELATIONS OPERATORS
	fmt.Println("----------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//END RELATION OPERATORS

	//LOGIC OPERATORS
	fmt.Println("----------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//END LOGIC OPERATORS

	//UNION OPERATORS
	fmt.Println("----------------------")
	number3 := 10
	number3++     // number3 = number3 + 1
	number3 += 15 //number3 = number3 + 15
	fmt.Println(number3)

	fmt.Println("----------------------")
	number3--     // number3 = number3 - 1
	number3 -= 15 //number3 = number3 +- 15
	fmt.Println(number3)

	fmt.Println("----------------------")
	number3 *= 3  // number3 = number3 * 3
	number3 /= 10 // number3 = number3 / 3
	number3 %= 3  // number3 = number3 % 3
	fmt.Println(number3)

	//END UNIONS OPERATORS
	fmt.Println("----------------------")
	var text string
	if number3 > 5 {
		text = "> 5"
	} else {
		text = "< 5"
	}
	fmt.Println(text)

}
