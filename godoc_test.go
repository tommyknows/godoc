package godoc

import "fmt"

// This is an example on how to use the GoDoc type
func ExampleGoDoc() {
	g := New("this is my documentation", "Ramon Ruettimann")
	fmt.Println(g)
	// Output: this is my documentation, written at [time] and authored by Ramon Ruettimann
}

// This is an example on how to use the GoDoc type with a label
func ExampleGoDoc_cool() {
	g := New("this is my cooler documentation", "Ramon Ruettimann")
	fmt.Println(g)
	// Output: this is my cooler documentation, written at [time] and authored by Ramon Ruettimann
}
