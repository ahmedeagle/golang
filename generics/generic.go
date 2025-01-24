package main

import "fmt"

// Filter selects elements from a slice that satisfy a predicate
func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Pair holds two values of any type
type Pair[A any, B any] struct {
	First  A
	Second B
}

func main() {
	// Example: Filter even numbers
	numbers := []int{1, 2, 3, 4, 5, 6}
	even := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(even) // Output: [2, 4, 6]

	// Example: Filter strings containing 'g'
	strings := []string{"go", "generics", "example"}
	filtered := Filter(strings, func(s string) bool {
		return len(s) > 4
	})
	fmt.Println(filtered) // Output: ["generics", "example"]

	/////////////////////////////////// generice struct ////////////////////////
	// Integer and string pair
	pair1 := Pair[int, string]{First: 1, Second: "One"}
	fmt.Printf("Pair 1: %v, %v\n", pair1.First, pair1.Second) // Output: 1, One

	// String and float pair
	pair2 := Pair[string, float64]{First: "Pi", Second: 3.14}
	fmt.Printf("Pair 2: %v, %v\n", pair2.First, pair2.Second) // Output: Pi, 3.14

	/////////////////////////////////
}
