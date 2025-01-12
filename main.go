package main

import (
	"fmt"

	"github.com/ahmedeagle/golang/oop"
)

// How to use interface in Go
type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 { // Rectangle implements Shape interface Area method
	return r.Width * r.Height
}

func (c Circle) Area() float64 { // Circle implements Shape interface Area method
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// How to use interface in Go
	var r oop.Shape = Rectangle{Width: 3, Height: 4}
	var c oop.Shape = Circle{Radius: 5}
	fmt.Println("Rectangle Area: ", r.Area())
	fmt.Println("Circle Area: ", c.Area())

	// How to achieve inhertance (emdedded class) in Go

	// Dog struct embedding Animal struct
	dog := oop.Dog{}
	dog.Speak()
	// Cat struct embedding Animal struct
	cat := oop.Cat{}
	cat.Speak()

	// How to use class-like methods in Go
	// Access package-level function (like static method)
	oop.StaticMethod()

	// Access class-like method
	oop.ClassMethod()

	// Create an instance of MyStruct and call an instance method
	ms := oop.MyStruct{InstanceAttr: "I am instance"}
	ms.InstanceMethod()

}
