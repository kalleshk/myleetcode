package main

import "fmt"

// Node represents a node in a singly linked list
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// InsertAtEnd inserts a node at the end of the list
func (l *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		// Traverse to the end of the list
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// InsertAtBeginning inserts a node at the beginning of the list
func (l *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
}

// Delete deletes the first node with the given value
func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the node to be deleted is the head
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	// Traverse the list to find the node before the one to be deleted
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	if current.Next == nil {
		fmt.Println("Value not found in the list")
		return
	}

	// Delete the node
	current.Next = current.Next.Next
}

// Display prints the elements of the linked list
func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	// Insert elements
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.InsertAtBeginning(5)

	fmt.Println("Linked List after insertion:")
	list.Display()

	// Delete an element
	list.Delete(20)
	fmt.Println("Linked List after deletion of 20:")
	list.Display()

	// Try to delete an element that doesn't exist
	list.Delete(100)
}
