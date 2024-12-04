package main

import (
	"fmt"
	"module/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from file main");
	auxiliar.Writer();
	
	erro := checkmail.ValidateFormat("camila@gmail.com");
	fmt.Println(erro);
}