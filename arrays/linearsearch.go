package arrays

func LinearSearch[T SliceValueTypes](list []T, search T) int {

	for i, element := range list {
		if element == search {
			return i
		}
	}

	return -1
}
