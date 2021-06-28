// TODO: build images and theory for each concept
// use this main to interact with each file
// Update readme with general problem and solution info
package main

import (
	"fmt"
	b "patterns/Builder"
	f "patterns/Factories"
	p "patterns/Prototype"
	s "patterns/SOLID"
)

func main() {
	s.DependencyInversion()
	b.BuilderFunctional()
	f.FactoryFunction()
	fmt.Println("Proto")
	p.CopySerial()
}
