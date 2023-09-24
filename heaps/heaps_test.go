package heaps

import (
	"fmt"
	"testing"

	"github.com/kreese89/dsa/heaps/generic_heap"
)

func TestBasicHeap(t *testing.T) {
	arr := []int{5, 7, -10, 6, 7, 2}
	sortedArr := []int{7, 7, 6, 5, 2, -10}
	heap := MakeHeap(arr)

	if heap.Empty() {
		t.Fatalf("Heap is empty!")
	}

	i := 0
	for !heap.Empty() {
		ele, err := heap.Pop()
		if err != nil {
			t.Fatalf("Popped on an empty heap")
		}
		if ele != sortedArr[i] {
			t.Fatalf("Expected: %v, Got: %v", sortedArr[i], ele)
		}
		i++
	}
}

func TestQueueLikeFunc(t *testing.T) {
	heap := NewHeap()
	err := heap.Push(1)

	if err != nil {
		t.Fatalf("Encountered error when pushing onto heap")
	}
	val, err := heap.Peek()
	if err != nil {
		t.Fatalf("Encountered error when peeking heap")
	} else if val != 1 {
		t.Fatalf("Encountered a different value")
	}

	heap.Push(5)
	val, err = heap.Peek()
	if err != nil {
		t.Fatalf("Encountered error when peeking heap")
	} else if val != 5 {
		t.Fatalf("Encountered %v, expected 5", val)
	}

	heap.Push(-1)
	if err != nil {
		t.Fatalf("Encountered error when peeking heap")
	} else if val != 5 {
		t.Fatalf("Encountered %v, expected 5", val)
	}

	heap.Pop()
	heap.Pop()
	heap.Pop()

	val, err = heap.Pop()

	if err == nil {
		t.Fatalf("Expected error on Pop on empty heap, got nil")
	}

}

func TestBasicHeapGeneric(t *testing.T) {
	arr := []int{5, 7, -10, 6, 7, 2}
	sortedArr := []int{7, 7, 6, 5, 2, -10}
	heap := generic_heap.MakeHeap(arr, true)

	if heap.Empty() {
		t.Fatalf("Heap is empty!")
	}
	fmt.Printf("%v", heap)

	i := 0
	for !heap.Empty() {
		ele, err := heap.Pop()
		if err != nil {
			t.Fatalf("Popped on an empty heap")
		}
		if ele != sortedArr[i] {
			t.Fatalf("Expected: %v, Got: %v", sortedArr[i], ele)
		}
		i++
	}
}
