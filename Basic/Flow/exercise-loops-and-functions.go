package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	var i int = 0
	for {
		n := (z*z - x) / (2 * z)
		if n < 0.0000001 && n > -0.0000001 {
			break
		} else {
			z -= n
		}
		i += 1
	}
	fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
