package main

import (
	"fmt"
)

// KeyValue is a struct for holding key-value pairs.
type KeyValue struct {
	key   string
	value int
}

// HashTable is a struct representing the hash table.
type HashTable struct {
	size  int
	table [][]KeyValue // 2D slice to handle collisions (chaining)
}

// NewHashTable initializes a hash table with a given size.
func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([][]KeyValue, size),
	}
}

// HashFunction generates an index for a given key.
func (ht *HashTable) HashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash + int(char)) % ht.size
	}
	return hash
}

// Insert adds a key-value pair into the hash table.
func (ht *HashTable) Insert(key string, value int) {
	index := ht.HashFunction(key)
	for i, kv := range ht.table[index] {
		if kv.key == key {
			// Update existing key with new value
			ht.table[index][i].value = value
			return
		}
	}
	// If key does not exist, append new key-value pair
	ht.table[index] = append(ht.table[index], KeyValue{key, value})
}

// Get retrieves the value for a given key.
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.HashFunction(key)
	for _, kv := range ht.table[index] {
		if kv.key == key {
			return kv.value, true
		}
	}
	return 0, false // Key not found
}

// Remove deletes a key-value pair from the hash table.
func (ht *HashTable) Remove(key string) bool {
	index := ht.HashFunction(key)
	for i, kv := range ht.table[index] {
		if kv.key == key {
			// Remove the key-value pair by re-slicing
			ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
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
