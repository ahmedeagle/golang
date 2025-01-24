package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

func (p *Person) updateStructName(name string) {

	p.Name = name
}

func main() {

	pers1 := Person{
		Name: "Emam",
	}

	//get instance of the struct and access all values
	fmt.Println(pers1.Name)

	//access struct method
	pers1.updateStructName("Not Emam")
	fmt.Println(pers1.Name)
}
