package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		dz := (z*z - x) / (2 * z)
		fmt.Printf("dz = %v\n", dz)
		z -= dz
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
