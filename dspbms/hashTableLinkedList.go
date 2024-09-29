package main

import (
	"fmt"
)

// Node represents an entry in the linked list
type Node struct {
	key   string
	value int
	next  *Node
}

// LinkedList represents the linked list for each bucket in the hash table
type LinkedList struct {
	head *Node
}

// HashTable is a struct representing the hash table, which uses an array of linked lists
type HashTable struct {
	size  int
	table []*LinkedList
}

// NewHashTable initializes a hash table with a given size
func NewHashTable(size int) *HashTable {
	// Create an array of linked lists (buckets)
	table := make([]*LinkedList, size)
	for i := range table {
		table[i] = &LinkedList{}
	}
	return &HashTable{
		size:  size,
		table: table,
	}
}

// HashFunction generates an index for a given key
func (ht *HashTable) HashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash + int(char)) % ht.size
	}
	return hash
}

// Insert adds a key-value pair into the hash table
func (ht *HashTable) Insert(key string, value int) {
	index := ht.HashFunction(key)
	linkedList := ht.table[index]

	// Check if the key already exists and update it
	for node := linkedList.head; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return
		}
	}

	// If key does not exist, insert a new node at the head of the linked list (prepend)
	newNode := &Node{key: key, value: value}
	newNode.next = linkedList.head
	linkedList.head = newNode
}

// Get retrieves the value for a given key
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.HashFunction(key)
	linkedList := ht.table[index]

	// Traverse the linked list to find the key
	for node := linkedList.head; node != nil; node = node.next {
		if node.key == key {
			return node.value, true
		}
	}
	return 0, false // Key not found
}

// Remove deletes a key-value pair from the hash table
func (ht *HashTable) Remove(key string) bool {
	index := ht.HashFunction(key)
	linkedList := ht.table[index]

	// Special case if the head node holds the key
	if linkedList.head != nil && linkedList.head.key == key {
		linkedList.head = linkedList.head.next // Remove the head
		return true
	}

	// Traverse the list to find the node before the one to remove
	for prev, curr := linkedList.head, linkedList.head.next; curr != nil; prev, curr = curr, curr.next {
		if curr.key == key {
			prev.next = curr.next // Remove the current node
			return true
		}
	}
	return false // Key not found
}

func main() {
	// Create a new hash table
	hashTable := NewHashTable(10)

	// Insert key-value pairs
	hashTable.Insert("apple", 100)
	hashTable.Insert("banana", 200)
	hashTable.Insert("orange", 300)

	// Retrieve values
	if value, found := hashTable.Get("apple"); found {
		fmt.Println("Value for 'apple':", value)
	} else {
		fmt.Println("'apple' not found")
	}

	// Update an existing key
	hashTable.Insert("apple", 150)
	if value, found := hashTable.Get("apple"); found {
		fmt.Println("Updated value for 'apple':", value)
	}

	// Remove a key
	if hashTable.Remove("banana") {
		fmt.Println("'banana' removed")
	}

	// Try to retrieve a removed key
	if value, found := hashTable.Get("banana"); found {
		fmt.Println("Value for 'banana':", value)
	} else {
		fmt.Println("'banana' not found")
	}
}
