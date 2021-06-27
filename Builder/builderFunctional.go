// Benefit of this method very easy to extend builder without aggregating nerw builders
// keeps a list of actions to perform on the object then peroform at once
package builder

import (
	"fmt"
)

type Person struct {
	name, position string
}

type personMod func(*Person)

type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAsA(job string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = job
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func BuinderFunctional() {
	b := PersonBuilder{}
	p := b.Called("Sam").WorksAsA("Developer").Build()
	fmt.Println(*p)
}
