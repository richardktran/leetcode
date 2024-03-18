package main

import (
	"fmt"
	"math"
)

func pivotInteger(n int) int {
	x := math.Sqrt((math.Pow(float64(n), 2) + float64(n)) / 2.0)

	y := int(x)

	if math.Abs(x-float64(y)) > 1e-06 {
		return -1
	}

	return y
}

func PivotIntegerTest() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
}

// func main() {
// 	PivotIntegerTest()
// }
