package prototype

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}
type Person struct {
	Name    string
	Address *Address
}

func DeepCopy() {

	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}

	// In this example Name will be copied but so will the pointer linking the the same address location to both
	// jane := john
	// jane.Name = "Jane"
	// jane.Address.StreetAddress = "13 Baker street"

	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane"
	jane.Address.StreetAddress = "13 Baker street"

	fmt.Print(john, john.Address)
	fmt.Print(jane, jane.Address)
}
