package prototype

import "fmt"

type AddressM struct {
	StreetAddress, City, Country string
}
type PersonM struct {
	Name    string
	Address *AddressM
	Friends []string
}

func (a *AddressM) DeepCopy() *AddressM {
	return &AddressM{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

func (p *PersonM) DeepCopy() *PersonM {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func CopyMethod() {
	john := PersonM{"John", &AddressM{"123 London Rd", "London", "UK"}, []string{"Chris", "Mary"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Print(john, john.Address)
	fmt.Print(jane, jane.Address)
}
