package merge_sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{5, 2, 7, 9, 8}
	sortedarr := []int{2, 5, 7, 8, 9}

	Sort(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedarr[i] {
			t.Fatalf("The array was not sorted correctly. Index: %v, expected: %v, got: %v", i, sortedarr[i], arr[i])
		}
	}

	arr = []int{-1, -1, -600, 6, 20}
	sortedarr = []int{-600, -1, -1, 6, 20}

	Sort(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedarr[i] {
			t.Fatalf("The array was not sorted correctly. Index: %v, expected: %v, got: %v", i, sortedarr[i], arr[i])
		}
	}
}
