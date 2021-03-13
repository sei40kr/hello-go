package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// len(a)=5
	a := make([]int, 5)
	printSlice("a", a)

	// len(b)=0 cap(b)=5
	b := make([]int, 0, 5)
	printslice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
