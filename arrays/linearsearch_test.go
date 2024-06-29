package arrays

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name,
		searchItem string
		list []string
		want int
	}{
		{
			"should be 2",
			"d",
			[]string{"a", "c", "d", "x", "k", "e"},
			2,
		},
		{
			"should be 4",
			"bk",
			[]string{"aa", "ac", "bd", "bx", "bk", "ae"},
			4,
		},
		{
			"should -1",
			"f",
			[]string{"a", "c", "d", "x", "k", "e"},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearSearch(tt.list, tt.searchItem)

			if result != tt.want {
				t.Errorf(" %v failed. %v is not equal to %v", tt.name, result, tt.want)
			}
		})
	}
}
