package main

import "fmt"

// Node represents a node in a binary tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// DFS Pre-order Traversal: Root -> Left -> Right
func DFS_PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value) // Visit the root
		DFS_PreOrder(node.left)       // Traverse left subtree
		DFS_PreOrder(node.right)      // Traverse right subtree
	}
}

// DFS In-order Traversal: Left -> Root -> Right
func DFS_InOrder(node *Node) {
	if node != nil {
		DFS_InOrder(node.left)        // Traverse left subtree
		fmt.Printf("%d ", node.value) // Visit the root
		DFS_InOrder(node.right)       // Traverse right subtree
	}
}

// DFS Post-order Traversal: Left -> Right -> Root
func DFS_PostOrder(node *Node) {
	if node != nil {
		DFS_PostOrder(node.left)      // Traverse left subtree
		DFS_PostOrder(node.right)     // Traverse right subtree
		fmt.Printf("%d ", node.value) // Visit the root
	}
}

// Main function to test the DFS operations
func main() {
	// Construct the following tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &Node{value: 1}
	root.left = &Node{value: 2}
	root.right = &Node{value: 3}
	root.left.left = &Node{value: 4}
	root.left.right = &Node{value: 5}

	// Perform Pre-order traversal
	fmt.Println("Pre-order traversal:")
	DFS_PreOrder(root) // Output: 1 2 4 5 3
	fmt.Println()

	// Perform In-order traversal
	fmt.Println("In-order traversal:")
	DFS_InOrder(root) // Output: 4 2 5 1 3
	fmt.Println()

	// Perform Post-order traversal
	fmt.Println("Post-order traversal:")
	DFS_PostOrder(root) // Output: 4 5 2 3 1
	fmt.Println()
}
