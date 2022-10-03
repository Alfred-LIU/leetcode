package DFS

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]
*/
func generateParenthesis(n int) []string {
	res := []string{}
	var backtrack func(path string, open int, close int)
	backtrack = func(path string, open, close int) {
		if close == n {
			res = append(res, path)
			return
		}
		if open < n {
			backtrack(path+"(", open+1, close)
		}
		if open > close {
			backtrack(path+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return res
}
