package main

import "fmt"

// Node represents a single node in the binary search tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// BST represents the binary search tree
type BST struct {
	root *Node
}

// Insert inserts a new value into the BST
func (bst *BST) Insert(value int) {
	bst.root = insertRec(bst.root, value)
}

// insertRec is a recursive helper function for inserting a new node
func insertRec(node *Node, value int) *Node {
	if node == nil {
		// If the node is nil, create a new Node and return it
		return &Node{value: value}
	}

	// Traverse the tree
	if value < node.value {
		// Insert in the left subtree
		node.left = insertRec(node.left, value)
	} else if value > node.value {
		// Insert in the right subtree
		node.right = insertRec(node.right, value)
	}

	return node
}

// Search searches for a value in the BST
func (bst *BST) Search(value int) bool {
	return searchRec(bst.root, value)
}

// searchRec is a recursive helper function for searching a value
func searchRec(node *Node, value int) bool {
	if node == nil {
		// If the node is nil, the value is not found
		return false
	}

	// Compare the value with the current node
	if value == node.value {
		return true
	} else if value < node.value {
		// Search in the left subtree
		return searchRec(node.left, value)
	} else {
		// Search in the right subtree
		return searchRec(node.right, value)
	}
}

// Delete deletes a value from the BST
func (bst *BST) Delete(value int) {
	bst.root = deleteRec(bst.root, value)
}

// deleteRec is a recursive helper function for deleting a node
func deleteRec(node *Node, value int) *Node {
	if node == nil {
		// If the node is nil, the value is not found
		return nil
	}

	// Traverse the tree to find the node to delete
	if value < node.value {
		node.left = deleteRec(node.left, value)
	} else if value > node.value {
		node.right = deleteRec(node.right, value)
	} else {
		// Node with the value found
		if node.left == nil {
			// Node has no left child, replace with right child
			return node.right
		} else if node.right == nil {
			// Node has no right child, replace with left child
			return node.left
		}

		// Node has two children, find the inorder successor (smallest in the right subtree)
		minRight := findMin(node.right)
		node.value = minRight.value

		// Delete the inorder successor
		node.right = deleteRec(node.right, minRight.value)
	}

	return node
}

// findMin finds the node with the minimum value in the BST (used in deletion)
func findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// InOrder traversal of the BST (for printing the tree in sorted order)
func (bst *BST) InOrder() {
	inOrderRec(bst.root)
	fmt.Println()
}

// inOrderRec is a recursive helper function for in-order traversal
func inOrderRec(node *Node) {
	if node != nil {
		inOrderRec(node.left)
		fmt.Printf("%d ", node.value)
		inOrderRec(node.right)
	}
}

func main() {
	// Create a new BST
	bst := &BST{}

	// Insert values into the BST
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	// Print the BST in order
	fmt.Println("In-order traversal after insertion:")
	bst.InOrder() // Output: 20 30 40 50 60 70 80

	// Search for values in the BST
	fmt.Println("Search for 40:", bst.Search(40)) // Output: true
	fmt.Println("Search for 25:", bst.Search(25)) // Output: false

	// Delete values from the BST
	bst.Delete(20)
	fmt.Println("In-order traversal after deleting 20:")
	bst.InOrder() // Output: 30 40 50 60 70 80

	bst.Delete(30)
	fmt.Println("In-order traversal after deleting 30:")
	bst.InOrder() // Output: 40 50 60 70 80

	bst.Delete(50)
	fmt.Println("In-order traversal after deleting 50:")
	bst.InOrder() // Output: 40 60 70 80
}
