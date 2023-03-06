package DFS

func isMatch(s string, p string) bool {
	lens := len(s)
	lenp := len(p)

	dp := make([][]bool, lens+1)

	for i := range dp {
		dp[i] = make([]bool, lenp+1)
	}

	dp[0][0] = true

	for j := 2; j < lenp+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	used := make(map[int]struct{})

	for i := 1; i < lens+1; i++ {
		for j := 1; j < lenp+1; j++ {
			sc := s[i-1]
			pc := p[j-1]

			if sc == pc || pc == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pc == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if p[j-2] == sc || p[j-2] == '.' {
					if _, ok := used[i*(lenp+1)+j]; !ok {
						dp[i][j] = dp[i-1][j]
						used[i*(lenp+1)+j] = struct{}{}
					}
				}
			}
		}
	}

	return dp[lens][lenp]
}
