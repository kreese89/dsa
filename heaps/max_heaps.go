package heaps

import (
	"errors"
	"fmt"
)

type MaxHeap struct {
	data []int
}

// Public Functions

func (heap *MaxHeap) Push(ele int) error {
	heap.data = append(heap.data, ele)
	return heap.updateVal(len(heap.data)-1, ele)
}

// constructs a new MaxHeap from a given array
// note that it does not heapify the input array but rather
// constructs a separate data structure using the initial data
func MakeHeap(arr []int) MaxHeap {
	data := make([]int, len(arr), cap(arr))
	copy(data, arr)
	heap := MaxHeap{data}

	var mid int
	mid = len(data) / 2
	for i := mid; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

func NewHeap() MaxHeap {
	heap := MaxHeap{[]int{}}

	return heap
}

func (heap *MaxHeap) ToArr() []int {
	var ret []int
	copy(ret, heap.data)
	return ret
}

func (heap *MaxHeap) Peek() (int, error) {
	if len(heap.data) == 0 {
		return -1, errors.New("Empty Heap")
	}
	return heap.data[0], nil
}

func (heap *MaxHeap) Pop() (int, error) {
	if len(heap.data) == 0 {
		return -1, errors.New("Empty heap")
	}
	top := heap.data[0]
	heap.data = heap.data[1:]
	heap.heapify(0)
	return top, nil
}

func (heap *MaxHeap) Replace(x int) {
}

func (heap *MaxHeap) Size() int {
	return len(heap.data)
}

func (heap *MaxHeap) Empty() bool {
	return len(heap.data) == 0
}

func HeapSort(arr []int) {
	// TODO
}

func (heap MaxHeap) String() string {
	return fmt.Sprintf("%v", heap.data)
}

// Internal Functions

// heapify is used internally to "balance" the heap and ensure
// that the heap property is maintained
// it is assumed that the left and right subtrees are already themselves heaped
func (heap *MaxHeap) heapify(i int) {
	left := heap.leftChild(i)
	right := heap.rightChild(i)
	var largest int
	if left < len(heap.data) && heap.data[left] > heap.data[i] {
		largest = left
	} else {
		largest = i
	}

	if right < len(heap.data) && heap.data[right] > heap.data[largest] {
		largest = right
	}

	if largest != i {
		tmp := heap.data[i]
		heap.data[i] = heap.data[largest]
		heap.data[largest] = tmp
		// since it is assumed that left and right were already heaps, we don't
		// have to rebalance the other subtree, just the subtree that we
		// just swapped with
		heap.heapify(largest)
	}

}

func (heap *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (heap *MaxHeap) rightChild(i int) int {
	return 2 * (i + 1)
}

func (heap *MaxHeap) parent(i int) int {
	var parInd int
	parInd = i / 2
	return parInd
}

func (heap *MaxHeap) updateVal(i int, newVal int) error {
	if i >= len(heap.data) {
		return errors.New("Index Out-of-bounds")
	}

	heap.data[i] = newVal
	parentInd := heap.parent(i)
	for i >= 0 && heap.data[parentInd] < heap.data[i] {
		tmp := heap.data[parentInd]
		heap.data[parentInd] = heap.data[i]
		heap.data[i] = tmp
		i = parentInd
	}

	return nil
}

// func (heap *MaxHeap) heapify(i int) {

// }
