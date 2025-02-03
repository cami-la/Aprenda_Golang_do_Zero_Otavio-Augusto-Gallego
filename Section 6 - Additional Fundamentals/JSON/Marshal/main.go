package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	d := Dog{"Rex", "Dálmata", 3}
	fmt.Println(d)

	dInJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(dInJSON))
	fmt.Println(bytes.NewBuffer(dInJSON))

	d2 := map[string]string{
		"name": "Toby",
		"race": "Poodle",
	}

	d2InJSON, err := json.Marshal(d2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(d2InJSON))
}
