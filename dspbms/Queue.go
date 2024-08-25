package main

import (
	"fmt"
)

// Queue represents a basic queue structure.
type Queue struct {
	items []interface{}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element at the front of the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Printf("Queue size: %d\n", queue.Size())

	fmt.Printf("Dequeue: %v\n", queue.Dequeue())
	fmt.Printf("Dequeue: %v\n", queue.Dequeue())

	fmt.Printf("Queue size after dequeue: %d\n", queue.Size())

	fmt.Printf("Is queue empty? %v\n", queue.IsEmpty())
}
