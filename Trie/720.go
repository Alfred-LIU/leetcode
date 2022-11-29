package Trie

/*
Given an array of strings words representing an English Dictionary, return the longest word in words that can be built one character at a time by other words in words.

If there is more than one possible answer, return the longest word with the smallest lexicographical order. If there is no answer, return the empty string.

Note that the word should be built from left to right with each additional character being added to the end of a previous word.
*/
func longestWord(words []string) string {
	trie := newTrieNode()
	for _, s := range words {
		trie.insert(s)
	}
	return trie.find()
}

type trieNode struct {
	children map[int32]*trieNode
	value    string
	isEnd    bool
}

func newTrieNode() trieNode {
	return trieNode{
		children: make(map[int32]*trieNode),
		isEnd:    true,
	}
}

func (tr *trieNode) insert(word string) {
	it := tr
	for _, c := range word {
		if _, ok := it.children[c-'a']; !ok {
			it.children[c-'a'] = &trieNode{
				children: make(map[int32]*trieNode),
				isEnd:    false,
			}
		}
		it = it.children[c-'a']
	}
	it.isEnd = true
	it.value = word
}

func (tr *trieNode) find() string {
	res := ""
	queue := make([]*trieNode, 0)
	queue = append(queue, tr)

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.isEnd == false {
			continue
		}
		if len(cur.value) > len(res) {
			res = cur.value
		}
		if len(cur.value) == len(res) && cur.value < res {
			res = cur.value
		}
		children := cur.children
		for _, c := range children {
			queue = append(queue, c)
		}
	}

	return res
}
