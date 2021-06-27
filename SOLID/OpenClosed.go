package solid

import "fmt"

// Open Closed Principle
// Open for extension, closed for modification
// Enterprise pattern: Specification

type Color int

// iota used for enumerating, starting red = 0
const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

// Filter products by color
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// New requirement to filter products by size means adding a new function
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

// New requirement filter product by size and color
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Each new requirement included modifying the filter struct which violates the Open Closed Principle
// By implementing specification we can fix this

type Specification interface {
	IsSatisfied(p *Product) bool
}

// color specification
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// Size specifications
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// To make a spec work for multiple filters e.g. size and color
// We need to implement the Composite Principle
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func OpenClosed() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Now if we want to create a new filter we would just create a new specification that followed the Spec interface
	// This follows the Open Closed principle by making sure we dont have change existing code
	// OPen for extension as in you can extend the functionality of these methods
	// But closed for modifying the existing code
	fmt.Printf("green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Using the composite principle to enable joint specifications
	fmt.Printf("green and large products:\n")
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}

	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}

}
