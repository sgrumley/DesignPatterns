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
	sing "patterns/Singleton"
)

func main() {
	fmt.Println("-------------------------------SOLID-------------------------------")
	s.DependencyInversion()
	fmt.Println("\n-------------------------------Builder-----------------------------")
	b.BuilderFunctional()
	fmt.Println("\n-------------------------------Factory-----------------------------")
	f.FactoryFunction()
	fmt.Println("\n-------------------------------Prototype---------------------------")
	p.CopySerial()
	fmt.Println("\n-------------------------------Singleton---------------------------")
	sing.SingletonDriver()
}
