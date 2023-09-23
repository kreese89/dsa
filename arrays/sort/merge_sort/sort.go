package merge_sort

func Sort(arr []int) {
	mergesort(arr, 0, len(arr)-1)
}

func mergesort(arr []int, p int, r int) {
	if p >= r {
		return
	}
	var q int
	q = (p + r) / 2
	mergesort(arr, p, q)
	mergesort(arr, q+1, r)
	merge(arr, p, q, r)
}

func merge(arr []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1, n1)
	right := make([]int, n2, n2)

	for i := 0; i < len(left); i++ {
		left[i] = arr[p+i]
	}
	for i := 0; i < len(right); i++ {
		right[i] = arr[q+i+1]
	}

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		// have to check if we've exhausted all options of left or right first
		if i >= len(left) {
			arr[k] = right[j]
			j++
		} else if j >= len(right) {
			arr[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}
