package arrays

import (
	"fmt"
	"math"
)

func BinarySearch[T SliceValueTypes](list []T, search T) int {

	low, high := 0, len(list)

	for ok := true; ok; ok = low < high {

		mid := int(math.Floor(float64(low + (high-low)/2)))
		value := list[mid]
		// fmt.Printf()
		fmt.Printf("List %v; search: %v, high: %v; low: %v; mid: %v \n", value, search, high, low, mid)

		if value == search {
			return mid
		}

		if value > search {
			high = mid
		} else {
			low = mid + 1
		}

	}

	return -1
}
