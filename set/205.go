package set

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
*/

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	setS := make(map[byte]struct{})
	setT := make(map[byte]struct{})

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}

		setS[s[i]] = struct{}{}
		setT[t[i]] = struct{}{}
	}

	return len(setS) == len(setT)
}
