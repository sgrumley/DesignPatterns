package factories

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// logic based on creating the struct
		fmt.Println("Under age")
	}
	return &Person{name, age, 2}
}

func FactoryFunction() {
	p1 := Person{"John", 22, 2}
	p2 := NewPerson("John", 33)
	fmt.Println(p1, p2)
}
