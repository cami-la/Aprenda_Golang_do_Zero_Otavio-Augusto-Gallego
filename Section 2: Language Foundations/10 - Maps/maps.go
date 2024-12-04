package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "João",
		"lastName": "Pedro",
	}
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"firstName": "cami",
			"lastName":  "Pedro",
		},
		"contact": {
			"phone": "123456789",
			"email": "email@email.com",
		},
	}
	fmt.Println(user2["contact"]["email"], 123)
	//delete(user2, "contact")
	fmt.Println(user2)

	user2["contact"] = map[string]string{
		"address": "Rua 1",
	}
	fmt.Println(user2)
}
