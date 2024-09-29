package main

import "fmt"

// TrieNode represents a single node in the trie
type TrieNode struct {
	children map[rune]*TrieNode // Map of children nodes
	isEnd    bool               // True if this node marks the end of a word
}

// Trie represents the trie (prefix tree) itself
type Trie struct {
	root *TrieNode
}

// NewTrie creates and initializes a new trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// Insert inserts a word into the trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		// If the character is not already in the map, create a new TrieNode
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		// Move to the next node in the trie
		currentNode = currentNode.children[char]
	}
	// Mark the end of the word
	currentNode.isEnd = true
}

// Search checks if a word exists in the trie
func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, char := range word {
		// If the character doesn't exist, the word is not in the trie
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		// Move to the next node
		currentNode = currentNode.children[char]
	}
	// Return true only if the last character is marked as the end of a word
	return currentNode.isEnd
}

func main() {
	// Create a new trie
	trie := NewTrie()

	// Insert words into the trie
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("banana")

	// Search for words in the trie
	fmt.Println(trie.Search("apple"))  // Output: true
	fmt.Println(trie.Search("app"))    // Output: true
	fmt.Println(trie.Search("banana")) // Output: true
	fmt.Println(trie.Search("ban"))    // Output: false
	fmt.Println(trie.Search("grape"))  // Output: false
}
