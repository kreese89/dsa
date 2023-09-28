package stacks

import (
	"testing"
)

func TestBasicStack(t *testing.T) {
	stack := NewStack[int]()

	_, err := stack.Peek()
	if err == nil {
		t.Fatalf("Expected fail on peek on empty stack")
	}

	_, err = stack.Pop()
	if err == nil {
		t.Fatalf("Expected fail on pop on empty stack")
	}

	err = stack.Push(5)

	if val, _ := stack.Peek(); val != 5 {
		t.Fatalf("Expected 5, got %v\n", val)
	}

	stack.Push(2)

	if val, _ := stack.Peek(); val != 2 {
		t.Fatalf("Expected 2, got %v\n", val)
	}

	if val, _ := stack.Pop(); val != 2 {
		t.Fatalf("Expected 2, got %v\n", val)
	}
	if val, _ := stack.Pop(); val != 5 {
		t.Fatalf("Expected 5, got %v\n", val)
	}

	if !stack.Empty() {
		t.Fatal("Expected stack to be empty, got nonempty stack")
	}

}

func TestStringStack(t *testing.T) {
	stack := NewStack[string]()
	vals := []string{"blah", "world!", "hello"}

	stack.Push("hello")
	stack.Push("world!")
	stack.Push("blah")

	for i := 0; i < len(vals); i++ {
		if val, _ := stack.Pop(); vals[i] != val {
			t.Fatalf("Expected %v, got %v", vals[i], val)
		}
	}

	if !stack.Empty() {
		t.Fatalf("Stack somehow not empty")
	}
}
