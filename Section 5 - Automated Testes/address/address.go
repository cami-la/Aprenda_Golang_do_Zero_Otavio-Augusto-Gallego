package address

import (
	"strings"
)

func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	santizedAddress := strings.ToLower(address)

	firstWord := strings.Split(santizedAddress, " ")[0]

	isValidType := false

	for _, typeAddress := range validTypes {
		if typeAddress == firstWord {
			isValidType = true
		}
	}

	if isValidType {
		return strings.Title(firstWord)
	}

	return "Invalid type"
}
