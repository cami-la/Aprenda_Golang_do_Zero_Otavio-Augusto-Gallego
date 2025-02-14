package main

import (
	"encoding/json"
	"fmt"
)

type Dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	dInJSON := `{"name":"Rex","race":"Dálmata","age":3}`

	var d Dog

	if erro := json.Unmarshal([]byte(dInJSON), &d); erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(d)

	d2InJSON := `{"name":"Toby","race":"Poodle"}`

	d2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(d2InJSON), &d2); erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(d2)

}
