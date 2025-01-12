package oop

import "fmt"

//Go follows the composition-over-inheritance principle, and interfaces enable polymorphism without the need for multiple inheritance.

type Animal struct{} //base struct

func (a Animal) Speak() {
	fmt.Println("Animal Speaking")
}

type Dog struct { // Dog struct embedding Animal struct
	Animal
}

type Cat struct { // Cat struct embedding Animal struct
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Dog Speaking")
}

func (c Cat) Speak() {
	fmt.Println("Cat Speaking")
}
