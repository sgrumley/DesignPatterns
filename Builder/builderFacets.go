// Multiple builders
// e.g aggregating builders under a single builder
package builder

import (
	"fmt"
)

type Person struct {
	// address
	Address, Zip, City string
	// job
	CompanyName, Position string
	Income                int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func newPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	it.person.Address = address
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) withZip(zip string) *PersonAddressBuilder {
	it.person.Zip = zip
	return it
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	it.person.CompanyName = companyName
	return it
}

func (it *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	it.person.Position = position
	return it
}

func (it *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	it.person.Income = income
	return it
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func BuilderFacets() {
	pb := newPersonBuilder()
	pb.
		Lives().
		At("24 East Street").
		In("Australia").
		withZip("4000").
		Works().
		At("Google").
		AsA("Software Developer").
		Earning(100000)

	// lives and works switches builder and can be changed at any point
	// Lives().
	// At("24 East Street").
	// In("Australia").
	// Works().
	// At("Google").
	// AsA("Software Developer").
	// Earning(100000).
	// Lives().
	//withZip("4000")

	person := pb.Build()
	fmt.Println(person)
	// fmt.Println("ndf")
}
