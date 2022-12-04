package topologySort

/*
There is a new alien language that uses the English alphabet. However, the order among the letters is unknown to you.

You are given a list of strings words from the alien language's dictionary, where the strings in words are sorted lexicographically by the rules of this new language.

Return a string of the unique letters in the new alien language sorted in lexicographically increasing order by the new language's rules. If there is no solution, return "". If there are multiple solutions, return any of them.

A string s is lexicographically smaller than a string t if at the first letter where they differ, the letter in s comes before the letter in t in the alien language. If the first min(s.length, t.length) letters are the same, then s is smaller if and only if s.length < t.length.
*/
func alienOrder(words []string) string {
	//compare adjascent words to find orderings..

	var adjList = make(map[byte][]byte, 0)
	var inDegree = make(map[byte]int, 0)

	for i := 0; i < len(words); i++ {
		wrd := words[i]
		for j := range wrd {
			if _, ok := adjList[wrd[j]]; !ok {
				adjList[wrd[j]] = make([]byte, 0)
				inDegree[wrd[j]] = 0
			}
		}
	}

	for i := 1; i < len(words); i++ {
		prev, curr := words[i-1], words[i]
		for j := range prev {
			if j < len(curr) && curr[j] != prev[j] {
				// in from prev[j] - curr[j]
				// indeg(curr[j]++)
				// add to adjlist(prev[j])
				adjList[prev[j]] = append(adjList[prev[j]], curr[j])
				inDegree[curr[j]]++
				break
			}
			if j == len(curr) {
				return ""
			}
		}
	}
	var queue = []byte{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	var ret = []byte{}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			top := queue[0]
			ret = append(ret, top)
			queue = queue[1:]
			// reduce inDeg of children by 1
			for _, child := range adjList[top] {
				inDegree[child]--
				if inDegree[child] == 0 {
					queue = append(queue, child)
				}
			}
		}
	}
	if len(adjList) != len(ret) {
		return ""
	}
	return string(ret)
	//invalid would mean that there is a cycle somewhere in our graph
}
