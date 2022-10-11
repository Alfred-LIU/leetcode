package SlidingWindow

func minWindow(s string, t string) string {
	count := 0
	m := make(map[byte]int, 0)

	for i, _ := range t {
		m[t[i]]++
		count++
	}

	if count > len(s) {
		return ""
	}

	var itr = string(make([]byte, len(s)))
	start, end := 0, 0

	for end < len(s) {
		if v, ok := m[s[end]]; ok {
			if v > 0 {
				count--
			}
			m[s[end]]--
		}

		for count <= 0 {
			if len(itr) >= len(s[start:end+1]) {
				itr = s[start : end+1]
			}

			if _, ok := m[s[start]]; ok {
				m[s[start]]++
				if m[s[start]] > 0 {
					count++
				}
			}
			start++
		}
		end++
	}

	if itr == string(make([]byte, len(s))) {
		return ""
	}

	return itr
}
