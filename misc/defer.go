package main

// defer.go
// author: Seong Yong-ju ( @sei40kr )

import "fmt"

func main() {
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	test()
	test2()
	fmt.Println("world")
}

func test() {
	defer fmt.Println("test world")
}

func test2() {
	defer fmt.Println("test2 world")
}
