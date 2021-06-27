// Benefit of this method very easy to extend builder without aggregating nerw builders
// keeps a list of actions to perform on the object then peroform at once
package builder

import (
	"fmt"
)

type PersonF struct {
	name, position string
}

type personMod func(*PersonF)

type PersonBuilderF struct {
	actions []personMod
}

func (b *PersonBuilderF) CalledF(name string) *PersonBuilderF {
	b.actions = append(b.actions, func(p *PersonF) {
		p.name = name
	})
	return b
}

func (b *PersonBuilderF) WorksAsA(job string) *PersonBuilderF {
	b.actions = append(b.actions, func(p *PersonF) {
		p.position = job
	})
	return b
}

func (b *PersonBuilderF) BuildF() *PersonF {
	p := PersonF{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func BuilderFunctional() {
	b := PersonBuilderF{}
	p := b.CalledF("Sam").WorksAsA("Developer").BuildF()
	fmt.Println(*p)
}
