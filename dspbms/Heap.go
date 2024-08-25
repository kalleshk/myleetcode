package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements the heap.Interface for a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &MinHeap{2, 1, 5}
	heap.Init(h)

	heap.Push(h, 3)
	fmt.Printf("Min-heap: %v\n", h)

	fmt.Printf("Pop minimum element: %d\n", heap.Pop(h))
	fmt.Printf("Min-heap after pop: %v\n", h)
}
