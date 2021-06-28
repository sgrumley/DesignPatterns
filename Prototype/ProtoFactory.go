package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type AddressF struct {
	StreetAddress, City string
	Suite               int
}
type Employee struct {
	Name   string
	Office *AddressF
}

var mainOffice = Employee{
	"", &AddressF{"123 East Dr", "London", 12},
}

var offSiteOffice = Employee{
	"", &AddressF{"321 West Dr", "London", 12},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewOffSiteOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&offSiteOffice, name, suite)
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)

	return &result
}

func ProtoFactory() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewOffSiteOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
