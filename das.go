package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64 = 2.0, 1119.0

	var x, y int = 2, 1119

	mid1 := int(math.Floor(a + (b-a)/2))
	fmt.Println(mid1)

	mid2 := int(math.Floor(float64(x + (y-x)/2)))
	fmt.Println(mid2)
}
