package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type AddressS struct {
	StreetAddress, City, Country string
}
type PersonS struct {
	Name    string
	Address *AddressS
	Friends []string
}

func (p *PersonS) DeepCopy() *PersonS {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := PersonS{}
	_ = d.Decode(&result)

	return &result
}

func CopySerial() {
	john := PersonS{"John", &AddressS{"123 London Rd", "London", "UK"}, []string{"Chris", "Mary"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
