package arrays

// sorts a give slice
func BubbleSort[T SliceValueTypes](list []T) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-i-1; j++ {
			curr := list[j]
			next := list[j+1]

			if curr > next {
				list[j+1] = curr
				list[j] = next
			}
		}
	}
}
