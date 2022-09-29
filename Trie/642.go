package Trie

import "sort"

type Node struct {
	//this 27 is important
	Children [27]*Node
	Count    int
	Word     string
}

type AutocompleteSystem struct {
	Root     *Node
	Curr     *Node
	CurrWord string
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	acs := AutocompleteSystem{
		Root: &Node{},
	}

	// Add Each word into Trie
	for i, word := range sentences {
		acs.Add(word, times[i])
	}
	acs.Curr = acs.Root

	return acs
}

func (this *AutocompleteSystem) Input(c byte) []string {
	var res []string

	//end of a sentence
	if c == '#' {
		this.Curr.Count++
		this.Curr.Word = this.CurrWord
		this.Curr = this.Root
		this.CurrWord = ""
		return res
	}

	this.CurrWord += string(c)
	idx := getIdx(c)

	//if no Node here
	if this.Curr.Children[idx] == nil {
		this.Curr.Children[idx] = &Node{}
		this.Curr = this.Curr.Children[idx]
		return res
	}

	this.Curr = this.Curr.Children[idx]
	list := make([]*Node, 0)
	getTopSearch(this.Curr, &list)

	sort.Slice(list, func(i, j int) bool {
		if list[i].Count == list[j].Count {
			return list[i].Word < list[j].Word
		}
		return list[i].Count > list[j].Count
	})

	for i := 0; i < 3 && i < len(list); i++ {
		res = append(res, list[i].Word)
	}

	return res
}

func (this *AutocompleteSystem) Add(word string, count int) {
	node := this.Root

	for _, c := range word {
		idx := getIdx(byte(c))

		if node.Children[idx] == nil {
			node.Children[idx] = &Node{}
		}
		node = node.Children[idx]
	}
	node.Count = count
	node.Word = word
}

func getTopSearch(node *Node, res *[]*Node) {
	if node.Count != 0 {
		*res = append(*res, node)
	}

	for _, n := range node.Children {
		if n != nil {
			getTopSearch(n, res)
		}
	}
}

func getIdx(c byte) int {
	if c == ' ' {
		return 26
	}
	return int(c - 'a')
}
