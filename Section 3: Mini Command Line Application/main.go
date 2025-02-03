package main

import (
	"Section_3_-_Mini_Command_Line_Application/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	application := app.Gerar()

	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
