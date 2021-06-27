package factories

import "fmt"

type EmployeeP struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("Unsupported role")
	}
}

func FactoryPrototype() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
