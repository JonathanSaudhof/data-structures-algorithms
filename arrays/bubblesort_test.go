package arrays

import (
	"slices"
	"testing"
)

func TestBubbleSortWithIntegersShouldReturnSorted(t *testing.T) {
	tests := []struct {
		name        string
		input, want []int
	}{
		{
			"should be 4, 5, 6, 7, 8, 10",
			[]int{5, 4, 6, 7, 8, 10},
			[]int{4, 5, 6, 7, 8, 10},
		},
		{
			"should be 4, 5, 6, 7, 8, 8, 10, 99",
			[]int{99, 8, 5, 4, 6, 7, 8, 10},
			[]int{4, 5, 6, 7, 8, 8, 10, 99},
		},
		{
			"should be 4, 5, 6, 7, 8, 8, 10, 99",
			[]int{99, 15, 12, 10, 9, 8, 7, 6},
			[]int{6, 7, 8, 9, 10, 12, 15, 99},
		},
		{
			"should be 1",
			[]int{1},
			[]int{1},
		},
		{
			"should be empty",
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.input)

			if !slices.Equal(tt.input, tt.want) {
				t.Errorf("%v: %v is not equal to %v", tt.name, tt.input, tt.want)
			}
		})
	}
}

func TestBubbleSortWithStringsShouldReturnSorted(t *testing.T) {
	tests := []struct {
		name        string
		input, want []string
	}{
		{
			"should be a c d e k x",
			[]string{"a", "c", "d", "x", "k", "e"},
			[]string{"a", "c", "d", "e", "k", "x"},
		},
		{
			"should be aa, ac, ae, bd, bk, bx",
			[]string{"aa", "ac", "bd", "bx", "bk", "ae"},
			[]string{"aa", "ac", "ae", "bd", "bk", "bx"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.input)

			if !slices.Equal(tt.input, tt.want) {
				t.Errorf("not equal")
			}
		})
	}
}
