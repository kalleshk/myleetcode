package main

import (
	"fmt"
)

// Stack represents a basic stack structure.
type Stack struct {
	items []interface{}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the element at the top of the stack.
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// Peek returns the element at the top of the stack without removing it.
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Printf("Stack size: %d\n", stack.Size())

	fmt.Printf("Pop: %v\n", stack.Pop())
	fmt.Printf("Pop: %v\n", stack.Pop())

	fmt.Printf("Stack size after pop: %d\n", stack.Size())

	fmt.Printf("Peek: %v\n", stack.Peek())

	fmt.Printf("Is stack empty? %v\n", stack.IsEmpty())
}
