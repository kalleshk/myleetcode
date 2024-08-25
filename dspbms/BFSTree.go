package main

import (
	"fmt"
)

// Node represents a single node in the binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BFS performs a breadth-first search (level-order traversal) on the tree
func (n *Node) BFS() {
	if n == nil {
		return
	}

	// Initialize a queue with the root node
	queue := []*Node{n}

	// Process the queue until it's empty
	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]

		// Print the current node's value
		fmt.Printf("%d ", current.Value)

		// Enqueue the left child if it exists
		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		// Enqueue the right child if it exists
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func main() {
	// Create a binary tree
	root := &Node{Value: 1}
	root.Left = &Node{Value: 2}
	root.Right = &Node{Value: 3}
	root.Left.Left = &Node{Value: 4}
	root.Left.Right = &Node{Value: 5}
	root.Right.Left = &Node{Value: 6}
	root.Right.Right = &Node{Value: 7}

	// Perform BFS traversal
	fmt.Println("BFS Traversal of the binary tree:")
	root.BFS()
	fmt.Println()
}
