package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// A slice just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding
	// elements of its underlying array.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
