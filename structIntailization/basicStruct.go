package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {

	//array are fixed u cannot add more
	var arr1 = []int{1, 2, 4}
	arr2 := [3]int{3, 434, 5}
	arr3 := []int{}

	arr4 := [5]int{1: 10, 2: 40}
	fmt.Println(arr4)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	// slice like array but more flexable
	//create slice form array
	arr5 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr5[2:4]
	fmt.Println(myslice)

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	//// struct

	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	fmt.Println(pers1)
	fmt.Println(pers1)
}
