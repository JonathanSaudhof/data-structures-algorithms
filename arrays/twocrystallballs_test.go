package arrays

import (
	"testing"
)

// having two crystal balls and an array of data
// representing if the crystal ball would brake
// the function returns the first position when it brakes
// [true, true, true, true, false, false]
func TestTwoCrystalBallsLinear(t *testing.T) {

	tests := []struct {
		name  string
		input []bool
		want  int
	}{
		{
			"should be -1",
			[]bool{false, false, false, false, false},
			-1,
		},
		{
			"should be 5",
			[]bool{false, false, false, false, false, true},
			5,
		},
		{
			"should 1",
			[]bool{false, true, true, true, true},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := TwoCrystalBallsLinear(tt.input)

			if idx != tt.want {
				t.Errorf("%v: %v is not equal to %v", tt.name, idx, tt.want)
			}
		})
	}

}

func TestTwoCrystalBallsSqrtN(t *testing.T) {

	tests := []struct {
		name  string
		input []bool
		want  int
	}{
		{
			"should be -1",
			[]bool{false, false, false, false, false},
			-1,
		},
		{
			"should be 5",
			[]bool{false, false, false, false, false, true},
			5,
		},
		{
			"should 1",
			[]bool{false, true, true, true, true},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := TwoCrystalBallsSqrtN(tt.input)

			if idx != tt.want {
				t.Errorf("%v: %v is not equal to %v", tt.name, idx, tt.want)
			}
		})
	}

}
