package generic_heap

import (
	"cmp"
	"errors"
	"fmt"
)

// Public Functions/Types

type Heap[T cmp.Ordered] struct {
	data     []T
	maximize bool // if max = true, will maximize on the comp operator
}

func MakeHeap[T cmp.Ordered](arr []T, maximize bool) Heap[T] {
	data := make([]T, len(arr), cap(arr))
	copy(data, arr)
	heap := Heap[T]{data, maximize}

	var mid int
	mid = len(data) / 2
	for i := mid; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

// creates and returns a new, empty heap
func NewHeap[T cmp.Ordered](maximize bool) Heap[T] {
	heap := Heap[T]{[]T{}, maximize}
	return heap
}

func (heap *Heap[T]) Peek() (T, error) {
	if len(heap.data) == 0 {
		var zero T
		return zero, errors.New("Empty Heap")
	}
	return heap.data[0], nil
}

func (heap *Heap[T]) Pop() (T, error) {
	if len(heap.data) == 0 {
		var zero T
		return zero, errors.New("Empty heap")
	}
	top := heap.data[0]
	heap.data = heap.data[1:]
	heap.heapify(0)
	return top, nil
}

func (heap *Heap[T]) Push(val T) error {
	heap.data = append(heap.data, val)
	return heap.updateVal(len(heap.data)-1, val)
}

func (heap *Heap[T]) Size() int {
	return len(heap.data)
}

func (heap *Heap[T]) Empty() bool {
	return len(heap.data) == 0
}

func (heap Heap[T]) String() string {
	return fmt.Sprintf("%v", heap.data)
}

// Private functions/types

func (heap *Heap[T]) heapify(i int) {
	left := heap.left(i)
	right := heap.right(i)
	optimum := i

	if left < len(heap.data) {
		if cmp.Less(heap.data[i], heap.data[left]) {
			if heap.maximize {
				optimum = left
			}
		} else {
			if !heap.maximize {
				optimum = left
			}
		}
	}

	if right < len(heap.data) {
		if cmp.Less(heap.data[optimum], heap.data[right]) {
			if heap.maximize {
				optimum = right
			}
		} else {
			if !heap.maximize {
				optimum = right
			}
		}
	}

	if optimum != i {
		tmp := heap.data[i]
		heap.data[i] = heap.data[optimum]
		heap.data[optimum] = tmp

		heap.heapify(optimum)
	}
}

func (heap *Heap[T]) left(i int) int {
	return 2*i + 1
}

func (heap *Heap[T]) updateVal(i int, newVal T) error {
	if i >= len(heap.data) {
		return errors.New("Index Out-of-bounds")
	}

	heap.data[i] = newVal
	parentInd := heap.parent(i)
	for i >= 0 {
		// need to break early depending on if we're maximizing or minimizing
		if heap.data[parentInd] >= heap.data[i] && heap.maximize {
			break
		} else if heap.data[parentInd] <= heap.data[i] && !heap.maximize {
			break
		}
		tmp := heap.data[parentInd]
		heap.data[parentInd] = heap.data[i]
		heap.data[i] = tmp
		i = parentInd
	}
	return nil
}

func (heap *Heap[T]) right(i int) int {
	return 2 * (i + 1)
}

func (heap *Heap[T]) parent(i int) int {
	var j int
	j = (i - 1) / 2
	return j
}
