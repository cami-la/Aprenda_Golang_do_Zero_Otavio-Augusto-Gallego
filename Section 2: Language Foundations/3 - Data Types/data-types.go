package main

import (
	"errors"
	"fmt"
)

func main () {
	//int8, int16, int32, int64

	var number int16 = 100;
	var number2 int16 = 100;

	fmt.Println(number, number2);

	//uint8, uint16, uint32, uint64
	var number3 uint32 = 1000;
	fmt.Println(number3)

	//alias
	//INT32 = RUNE
	var number4 rune = 12456
	fmt.Println(number4)

	//BYTE = UINT8
	var number5 byte = 123
	fmt.Println(number5)

	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 1230000000.45
	fmt.Println(realNumber2)

	realNumber3 := 123456.67
	fmt.Println(realNumber3)

	//END REAL NUMBERS

	var str string = "Text"
	fmt.Println(str)

	str2 := "text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//END STRINGS

	var text string
	var int int16

	fmt.Println(text, int)

	var boolean1 bool
	var boolean2 bool = true

	fmt.Println(boolean1, boolean2)

	var erro error
	var erro2 error = errors.New("Internal Error")
	fmt.Println(erro, erro2)
}