package string

import "strings"

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i, j := 0, len(s)-1

	for i < j {
		for (i < j) && !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z')) {
			i += 1
		}

		for (i < j) && !((s[j] >= '0' && s[j] <= '9') || (s[j] >= 'A' && s[j] <= 'Z')) {
			j -= 1
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
