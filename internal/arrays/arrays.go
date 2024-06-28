package arrays

func BubbleSort[T string | int](array []T) []T {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			curr := array[j]
			next := array[j+1]

			if curr > next {
				array[j+1] = curr
				array[j] = next
			}
		}
	}
	return array
}

type SliceValueTypes interface {
	string | int
}
