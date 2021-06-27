solid main

type Document interface {

}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {

}

func (m MultiFunctionPrinter) Print(d Document) {

}
func (m MultiFunctionPrinter) Fax(d Document) {

}
func (m MultiFunctionPrinter) Scan(d Document) {

}

// Even if the printer doesn't have all functionality the machine interface forces us to implement the functions
type OldFunctionPrinter struct {

}

func (m OldFunctionPrinter) Print(d Document) {

}
// Deprecated
func (o OldFunctionPrinter) Fax(d Document) {
	panic("Not supported")
}
// deprecated
func (o OldFunctionPrinter) Scan(d Document) {
	panic("Not supported")
}

// Interface Segregation Principle
// Avoid having too much in an interface

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// Only have print functionality
type MyPrinter struct {

}

func (m MyPrinter) Print(d Document){

}

type PhotoCopier struct {

}

func (p PhotoCopier) Scan(d Document){

}
func (p PhotoCopier) Print(d Document){

}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document){
	m.printer.Print(d)
}

