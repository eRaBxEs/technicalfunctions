package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	var x float64
	var y float64

	x = math.Sqrt((b * b) - (4 * a * c))

	x = (-b + x) / (2 * a)

	y = math.Sqrt((b * b) - (4 * a * c))

	y = (-b - y) / (2 * a)

	return x, y
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}
