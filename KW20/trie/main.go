package main

import (
	"fmt"
)

// TrieNode represents a single node in the Trie.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie represents the Trie structure with a root node.
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie.
func (t *Trie) Insert(word string) {

}

// Search checks if a word is in the Trie.
func (t *Trie) Search(word string) bool {
	return false
}

// StartsWith checks if any word in the Trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return false
}

// GetWordsWithPrefix returns all words in the Trie that start with the given prefix.
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	return []string{}
}

// collectWords is a helper function to perform a search and collect words from the given node.
func (t *Trie) collectWords(node *TrieNode, currentWord string, words *[]string) {

}

func main() {
	trie := NewTrie()
	words := []string{"apple", "app", "apricot", "banana", "band", "bandana", "cat", "car", "cart"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Testing Search and StartsWith
	fmt.Println(trie.Search("apple"))   // Output: true
	fmt.Println(trie.Search("app"))     // Output: true
	fmt.Println(trie.Search("apricot")) // Output: true
	fmt.Println(trie.Search("banana"))  // Output: true
	fmt.Println(trie.Search("bananas")) // Output: false
	fmt.Println(trie.StartsWith("ap"))  // Output: true
	fmt.Println(trie.StartsWith("ba"))  // Output: true
	fmt.Println(trie.StartsWith("ca"))  // Output: true
	fmt.Println(trie.StartsWith("da"))  // Output: false

	// Testing GetWordsWithPrefix
	fmt.Println(trie.GetWordsWithPrefix("ap"))  // Output: [apple app apricot]
	fmt.Println(trie.GetWordsWithPrefix("ba"))  // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("ca"))  // Output: [cat car cart]
	fmt.Println(trie.GetWordsWithPrefix("ban")) // Output: [banana band bandana]
	fmt.Println(trie.GetWordsWithPrefix("xyz")) // Output: []
}
