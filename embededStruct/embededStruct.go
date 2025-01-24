package main

import "fmt"

// Base struct
type Address struct {
	City    string
	Country string
}

// Outer struct embedding Address
type Person struct {
	Name    string
	Age     int
	Address // Embedded struct
}

// Outer struct with an explicit method
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s from %s, %s.\n", p.Name, p.City, p.Country)
}

func main() {
	// Create a Person with embedded Address
	person := Person{
		Name: "Ahmed",
		Age:  30,
		Address: Address{
			City:    "Dubai",
			Country: "UAE",
		},
	}

	// Access fields and methods directly
	fmt.Println(person.Name) // Output: Ahmed
	fmt.Println(person.City) // Output: Dubai (from embedded Address)
	person.Greet()           // Output: Hello, my name is Ahmed from Dubai, UAE.
}
