package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name,
		searchItem string
		list []string
		want int
	}{
		{
			"should be 2",
			"d",
			[]string{"a", "c", "d", "e", "f", "g"},
			2,
		},
		{
			"should be 7",
			"cg",
			[]string{"aa", "ac", "ae", "bd", "bk", "bx", "cd", "cg"},
			7,
		},
		{
			"should -1",
			"f",
			[]string{"a", "c", "d", "e", "k", "x"},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.list, tt.searchItem)

			if result != tt.want {
				t.Errorf(" %v failed. %v is not equal to %v", tt.name, result, tt.want)
			}
		})
	}
}
