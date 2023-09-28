package stacks

import (
	"errors"
)

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() Stack[T] {
	stack := Stack[T]{[]T{}}
	return stack
}

func (stack *Stack[T]) Empty() bool {
	if len(stack.data) == 0 {
		return true
	}
	return false
}

func (stack *Stack[T]) Peek() (T, error) {
	var zero T
	if stack.Empty() {
		return zero, errors.New("Tried peeking on an empty stack")
	}

	return stack.data[0], nil
}

func (stack *Stack[T]) Pop() (T, error) {
	var zero T
	if stack.Empty() {
		return zero, errors.New("Tried popping on an empty stack")
	}
	top := stack.data[0]
	stack.data = stack.data[1:]
	return top, nil
}

func (stack *Stack[T]) Push(ele T) error {
	stack.data = append([]T{ele}, stack.data...)
	return nil
}

func (stack *Stack[T]) Size() int {
	return len(stack.data)
}
