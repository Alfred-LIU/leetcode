package BFS

/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return all the shortest transformation sequences from beginWord to endWord, or an empty list if no such sequence exists. Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].



Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
Explanation: There are 2 shortest transformation sequences:
"hit" -> "hot" -> "dot" -> "dog" -> "cog"
"hit" -> "hot" -> "lot" -> "log" -> "cog"
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: []
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]struct{})
	current, trace := make(map[string]struct{}), make(map[string][]string)
	for _, s := range wordList {
		dict[s] = struct{}{}
		trace[s] = make([]string, 0)
	}

	dict[beginWord] = struct{}{}
	trace[beginWord] = make([]string, 0)
	current[beginWord] = struct{}{}

	_, ok := current[endWord]
	for len(current) != 0 && !ok {
		for word, _ := range current {
			delete(dict, word)
		}
		next := make(map[string]struct{})
		for word, _ := range current {
			for candidate, _ := range allAjacents(word, wordList) {
				if _, ok := dict[candidate]; ok {
					trace[candidate] = append(trace[candidate], word)
					next[candidate] = struct{}{}
				}
			}
		}
		current = next
		_, ok = current[endWord]
	}

	var results [][]string
	if len(current) != 0 {
		backtrace(&results, trace, []string{}, endWord, beginWord)
	}
	return results
}

func backtrace(results *[][]string, trace map[string][]string, path []string, word string, begin string) {
	if word == begin {
		*results = append(*results, append([]string{word}, path...))
	} else {
		for _, prev := range trace[word] {
			backtrace(results, trace, append([]string{word}, path...), prev, begin)
		}
	}
}

func allAjacents(word string, wordList []string) map[string]bool {
	res := make(map[string]bool, 0)
	for _, w := range wordList {
		if isAdjacent(word, w) {
			res[w] = true
		}
	}

	return res
}

func isAdjacent(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	count := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			count++
		}
	}

	return count == 1
}
