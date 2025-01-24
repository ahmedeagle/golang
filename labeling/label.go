package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			println(i, j)
			if j == 1 {
				goto nextCode
			}
		}
	}

nextCode:
	fmt.Println("executing next code after goto")
}
