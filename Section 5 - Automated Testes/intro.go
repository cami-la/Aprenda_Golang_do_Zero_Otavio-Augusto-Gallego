package main

import (
	"Section_5_-_Automated_Testes/address"
	"fmt"
)

func main() {
	addressType := address.AddressType("rua Nestor Silva, 70")
	fmt.Println(addressType)
}
