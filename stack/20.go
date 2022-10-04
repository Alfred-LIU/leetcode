package stack

/*
20. Valid Parentheses
Easy

15914

796

Add to List

Share
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
*/

func isValid(s string) bool {
	m := map[int32]int32{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]int32, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' || v == '}' || v == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != m[v] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}
