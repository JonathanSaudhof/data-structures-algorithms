package arrays

import "math"

type SliceValueTypes interface {
	string | int
}

// given a as the lower, and b as the higher end
// the function calculates the next highest integer midpoint
func calculateMidpoint(a, b int) int {
	return int(math.Floor(float64(a + (b-a)/2)))
}
