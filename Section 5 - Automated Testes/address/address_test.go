package address

import "testing"

type fakeAddress struct {
	receivedAddress string
	expectedType    string
}

func TestAddressType(t *testing.T) {
	fakeAddress := []fakeAddress{
		{"Rua Esmeraldino Bandeira", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada dos Bandeirantes", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalid type"},
		{"", "Invalid type"},
	}

	for _, fakeAddress := range fakeAddress {
		receivedType := AddressType(fakeAddress.receivedAddress)
		if receivedType != fakeAddress.expectedType {
			t.Errorf("Received %s, expected %s", receivedType, fakeAddress.expectedType)
		}
	}
}
