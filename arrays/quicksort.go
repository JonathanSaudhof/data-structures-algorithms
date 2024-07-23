package arrays

func sort(arr *[]int, low int, high int) {
	if low > high {
		return
	}

	pivot := pivot(arr, low, high)
	sort(arr, low, pivot-1)
	sort(arr, pivot+1, high)
}

func pivot(arr *[]int, low int, high int) int {
	pivotElement := (*arr)[high]

	idx := low - 1

	for i, j := low, high; i < j; i++ {
		if (*arr)[i] < pivotElement {
			idx++
			swap(arr, idx, i)
		}
		if (*arr)[j] > pivotElement {
			j--
		}
	}
	idx++
	swap(arr, idx, high)

	return idx
}

func swap(arr *[]int, idxA int, idxB int) {
	tmp := (*arr)[idxA]
	(*arr)[idxA] = (*arr)[idxB]
	(*arr)[idxB] = tmp
}

func QuickSort(arr *[]int) {
	var low, high int = 0, len(*arr) - 1

	sort(arr, low, high)
}
