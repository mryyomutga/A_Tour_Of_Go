package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	lz, z := x, float64(1)
	for math.Abs(z-lz) >= 1e-6 {
		lz, z = z, z - (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
