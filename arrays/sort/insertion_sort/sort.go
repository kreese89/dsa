package insertion_sort

// Implementation of the insertion sort algorithm
func Sort(arr []int) {
	// at each iteration i=k, the subarray
	// arr[0, ..., k-1] will be sorted
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
