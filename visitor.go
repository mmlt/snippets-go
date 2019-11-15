package main

import "fmt"

// Visitor pattern
// The visitor knows the element types. The elements only know that the visitor implements the ElementVisitor interface
// so multiple visitors implementations can be used.

// ElementVisitor defines what methods a visitor for Elements must implement.
type ElementVisitor interface {
	ElementA(e ElementA)
	ElementB(e ElementB)
}


// Define 2 objects that can be visited.

type ElementA string

func (e ElementA) Accept(v ElementVisitor) {
	v.ElementA(e)
}

type ElementB int

func (e ElementB) Accept(v ElementVisitor) {
	v.ElementB(e)
}


// Define a Visitor that simple print elements.
type SimpleElementVisitor string

func (v SimpleElementVisitor) ElementA(e ElementA) {
	fmt.Println(v, "element A says", e)
}

func (v SimpleElementVisitor) ElementB(e ElementB) {
	fmt.Println(v, "element B says", e)
}

// Define a Visitor that print type info.
type DetailElementVisitor string

func (v DetailElementVisitor) ElementA(e ElementA) {
	fmt.Printf("%s element A type=%T value=%v\n", v, e, e)
}

func (v DetailElementVisitor) ElementB(e ElementB) {
	fmt.Printf("%s element B type=%T value=%v\n", v, e, e)
}

func main() {
	a := ElementA("one")
	b := ElementB(2)

	sv := SimpleElementVisitor("Simple")
	a.Accept(sv)
	b.Accept(sv)

	dv := DetailElementVisitor("Detailed")
	a.Accept(dv)
	b.Accept(dv)
}