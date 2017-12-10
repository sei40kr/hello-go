package main

// panic.go
// author: Seong Yong-ju ( @sei40kr )

import "fmt"

func main() {
	test()
	fmt.Println("world")
}

func test() {
	defer func(msg string) {
		if error := recover(); error != nil {
			fmt.Println(msg)
			fmt.Println(error)
			fmt.Println("we recovered")
		}
	}("hello")
	panic("panic happened")
}
