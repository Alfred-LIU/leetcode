package Trie

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/
type TrieNode struct {
	isEnd    bool
	children [26]*TrieNode
}

type Trie struct {
	root TrieNode
}

func Constructor() Trie {
	return Trie{
		root: TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	next := &this.root
	for _, c := range word {
		if next.children[c-'a'] == nil {
			next.children[c-'a'] = &TrieNode{}
		}
		next = next.children[c-'a']
	}
	next.isEnd = true
}

func (this *Trie) Search(word string) bool {
	next := &this.root
	for _, c := range word {
		if next.children[c-'a'] == nil {
			return false
		}
		next = next.children[c-'a']
	}
	return next.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	next := &this.root
	for _, c := range prefix {
		if next.children[c-'a'] == nil {
			return false
		}
		next = next.children[c-'a']
	}
	return true
}
