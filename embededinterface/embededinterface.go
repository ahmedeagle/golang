package main

import "fmt"

// Define two interfaces
type Reader interface {
	Read() string
	ReadV2() string
}

type Writer interface {
	Write(data string)
}

// Embed Reader and Writer into a new interface
type ReadWriter interface {
	Reader
	Writer
}

// Struct implementing ReadWriter
type File struct{}

func (f File) Read() string {
	return "Reading from file"
}

func (f File) ReadV2() string {
	return "ReadingV2 from file"
}

func (f File) Write(data string) {
	fmt.Println("Writing to file:", data)
}

func main() {
	// File satisfies ReadWriter because it implements both Reader and Writer
	// we got all methods from different struct and now we can access all of them from only one struct
	var rw ReadWriter = File{}

	// Access methods from embedded interfaces
	fmt.Println(rw.Read())    // Output: Reading from file
	fmt.Println(rw.ReadV2())  // Output: ReadingV2 from file
	rw.Write("Hello, World!") // Output: Writing to file: Hello, World!
}
