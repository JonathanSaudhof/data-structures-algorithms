package arrays_test

import (
	"testing"

	"portfolio/das/arrays"
)

func TestQuickSort_ShouldBeSorted(t *testing.T) {
	unsortedArray := []int{
		5, 9, 3, 7, 2, 8, 1, 6, 4,
	}

	arrays.QuickSort(&unsortedArray)

	expected := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	if len(unsortedArray) != len(expected) {
		t.Errorf("Expected %v but got %v", expected, unsortedArray)
	}

	for i := 0; i < len(unsortedArray); i++ {
		if unsortedArray[i] != expected[i] {
			t.Errorf("Expected %v but got %v", expected, unsortedArray)
		}
	}
}
