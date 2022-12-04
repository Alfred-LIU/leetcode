package DFS

/*
Given a pattern and a string s, return true if s matches the pattern.

A string s matches a pattern if there is some bijective mapping of single characters to strings such that if each character in pattern is replaced by the string it maps to, then the resulting string is s. A bijective mapping means that no two characters map to the same string, and no character maps to two different strings.
*/

func wordPatternMatch(pattern string, s string) bool {
	if len(pattern) == 0 && len(s) == 0 {
		return true
	}
	if len(pattern) == 0 || len(s) == 0 {
		return false
	}

	m := make(map[byte]string)
	set := make(map[string]struct{})
	return dfs291(0, 0, pattern, s, m, set)
}

func dfs291(i, j int, pattern, s string, m map[byte]string, set map[string]struct{}) bool {
	if i == len(pattern) && j == len(s) {
		return true
	}
	if i == len(pattern) || j == len(s) {
		return false
	}

	p := pattern[i]
	for k := j + 1; k <= len(s); k++ {
		word := s[j:k]
		v, patternOK := m[p]
		_, setOK := set[word]
		if patternOK && word == v {
			if dfs291(i+1, k, pattern, s, m, set) {
				return true
			}
		} else if !patternOK && !setOK {
			m[p] = word
			set[word] = struct{}{}
			if dfs291(i+1, k, pattern, s, m, set) {
				return true
			}
			delete(m, p)
			delete(set, word)
		}
	}

	return false
}
