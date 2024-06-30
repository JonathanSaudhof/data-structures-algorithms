package arrays

import "math"

// having two crystal balls and an array of data
// representing if the crystal ball would brake
// the function returns the first position when it brakes
// [false, false, false, false, true, true]
// returns 4
func TwoCrystalBallsLinear(list []bool) int {
	var low, high int = 0, len(list)

	for ok := true; ok; {

		if low >= high {
			return -1
		}

		middle := calculateMidpoint(low, high)
		value := list[middle]

		if value == true {
			high = middle
			ok = false
		} else {
			low = middle + 1
		}

	}

	for i, value := range list[low : high+1] {
		if value == true {
			return i + low
		}
	}

	return -1
}

// having two crystal balls and an array of data
// representing if the crystal ball would brake
// the function returns the first position when it brakes
// in a less-linear runtime, sqrt(N)
// [false, false, false, false, true, true]
// returns 4
func TwoCrystalBallsSqrtN(list []bool) int {
	listLength := len(list)
	sqrt := math.Sqrt(float64(listLength))
	var jumpSize int = int(math.Floor(sqrt))
	idx := jumpSize

	for ; idx < listLength; idx += jumpSize {

		if list[idx] == true {
			break
		}

	}

	idx -= jumpSize

	for i, value := range list[idx:] {
		if value == true {
			return i + idx
		}
	}

	return -1
}
