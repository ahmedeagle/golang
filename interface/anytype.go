package main

import "fmt"

func showVal(t interface{}) {

	fmt.Println(t)
}

func main() {

	var value interface{} = 10
	showVal(value)

	var valuesrt interface{} = "string "
	showVal(valuesrt)
}
