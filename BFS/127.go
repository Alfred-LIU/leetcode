package BFS

/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.



Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//build map
	allwords := wordList
	allwords = append(allwords, beginWord)
	m := buildMap(allwords, wordList)

	adjacents := m[beginWord]
	seen := make(map[string]bool, len(wordList))
	changes := 1

	for len(adjacents) > 0 {
		changes++

		tmp := make(map[string]bool, 0)
		for w, _ := range adjacents {
			if w == endWord {
				return changes
			}
			if seen[w] {
				continue
			}
			seen[w] = true

			for k, _ := range m[w] {
				tmp[k] = true
			}
		}
		adjacents = tmp
	}

	return 0
}

func buildMap(allwords []string, wordList []string) map[string]map[string]bool {
	res := make(map[string]map[string]bool, 0)
	for _, v := range allwords {
		res[v] = allAjacents(v, wordList)
	}
	return res
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
