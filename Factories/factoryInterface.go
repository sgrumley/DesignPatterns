package factories

import "fmt"

type PersonInt interface {
	SayHello()
}

type personI struct {
	name string
	age  int
}

func (p *personI) SayHello() {
	fmt.Printf("Hi my name is %s, I am %d years old\n", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("I am too tired to say hello")
}

func NewPersonI(name string, age int) PersonInt {
	if age > 150 {
		return &tiredPerson{name, age}
	}
	return &personI{name, age}
}

func FactoryInterface() {
	p := NewPersonI("Jeff", 15)
	p.SayHello()

	t := NewPersonI("James", 155)
	t.SayHello()
}
