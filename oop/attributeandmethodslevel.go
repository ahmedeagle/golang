package oop

import "fmt"

var StaticAttr = "I am static" // Package-level variable (like static attribute)

type MyStruct struct {
	InstanceAttr string
}

// Static method equivalent (using a package-level function)
func StaticMethod() {
	fmt.Println("This is a static method.")
}

// Method associated with a struct (similar to class method)
func (ms MyStruct) InstanceMethod() {
	fmt.Println("This is an instance method.")
}

// Class-like method (function that works with the struct itself)
func ClassMethod() {
	fmt.Println("Class method, accessing static attribute:", StaticAttr)
}
