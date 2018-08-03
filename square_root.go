package main

import (
	"fmt"
	"math"
)

// Sqrt is a function that tries to imitate
// what math.Sqrt does
func Sqrt(x float64) float64 {

	z := 1.000000
	zPrevious := 0.0
	distance := 1.0
	minDistance := 0.000001
	i := 0

	for ; math.Abs(distance) > minDistance; i++ {
		//    for ; i<10; i++ {
		distance := (z*z - x) / (2 * z)
		z = z - distance
		if math.Abs(zPrevious-z) < minDistance {
			break
		}
		zPrevious = z
		//		fmt.Println(distance,z,i)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2.0))
	fmt.Println(math.Sqrt(2))
}
