// Motivation: Some objects are hard to construct builder provides an api

// Some objects are simple and can be created in a single constructor call
// Other objects have a lot of set up involved
// Having a factory function with 10 arguments is not productive

// OPt for piecewise (piece by piece) construction 
// Builder provides an API for constructing an object step by step

package main 

import (
	"fmt"
	"strings"
)

func antiBuilder() {
	// basic web server e.g. list to html
	// strings.Builder accepts strings into a buffer and concatinates them
	sb := strings.Builder{}
	words := []string{"hello", "world"}
	sb.Reset();
	//<ul><li>..</ul></li>
	// this is anti builder due to the amount of steps adding strings such as ul, li
	sb.WriteString("<ul>");
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>");
	fmt.Println(sb.String())

}

// Builder implementation
const indentSize = 2;

type HtmlElement struct {
	name, text string 
	elements []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0);
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent+1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

type HtmlBuilder struct {
	rootName string 
	root HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{ rootName, HtmlElement{rootName, "", []HtmlElement{} }}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{ childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// Bonus: fluent interfaces - allows chaining
// returning a pointer to the reciever
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{ childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func Builder() {
	b := NewHtmlBuilder("ul");
	b.AddChild("li", "hello");
	b.AddChild("li", "world");
	fmt.Println(b.String());
}

func fluentBuilder() {
	b := NewHtmlBuilder("ul");
	b.AddChildFluent("li", "hello").AddChildFluent("li", "world");
	fmt.Println(b.String());
}

func main() {
	Builder();
	antiBuilder();
	fluentBuilder();
}