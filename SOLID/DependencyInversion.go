package solid

import "fmt"

// High Level Modules (usually business logic) should not depend on Low Level Modules (typically system or storage)
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low level module (Basically the data)
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// High level module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationshipBrowser
}

// If relationships (Low level) decides to change it's storage type e.g. DB
// This code will break, so we create the interface relationship Browser
//func (r *Research) Investigate() {
//	relations := r.relationships.relations
//	for _, rel := range relations {
//		if rel.from.name == "John" && rel.relationship == Parent {
//			fmt.Println("John has a child called ", rel.to.name)
//		}
//	}
//}

// new function that depends on abstractions
func (r *Research) Investigate() {

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child named", p.name)
	}
}

func DependencyInversion() {
	parent := Person{"John"}
	child1 := Person{"Mathew"}
	child2 := Person{"Frank"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	//r := Research{relationships}
	r := Research{&relationships}
	r.Investigate()
}
