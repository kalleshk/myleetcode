package main

import (
	"fmt"
)

// Node represents a single node in the binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node with the given value to the tree
func (n *Node) Insert(value int) {
	if value < n.Value {
		// Insert in the left subtree
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		// Insert in the right subtree
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search looks for a value in the tree and returns true if found, false otherwise
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

// InOrder prints the values of the tree in sorted order (in-order traversal)
func (n *Node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrder()
}

// main demonstrates usage of the binary search tree
func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	fmt.Println("In-order traversal of the binary search tree:")
	root.InOrder()
	fmt.Println()

	// Searching for a value
	fmt.Println("Searching for value 7:", root.Search(7))
	fmt.Println("Searching for value 20:", root.Search(20))
}
