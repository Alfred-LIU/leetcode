package stack

import (
	"sort"
	"strconv"
)

/*
Given a string formula representing a chemical formula, return the count of each atom.

The atomic element always starts with an uppercase character, then zero or more lowercase letters, representing the name.

One or more digits representing that element's count may follow if the count is greater than 1. If the count is 1, no digits will follow.

For example, "H2O" and "H2O2" are possible, but "H1O2" is impossible.
Two formulas are concatenated together to produce another formula.

For example, "H2O2He3Mg4" is also a formula.
A formula placed in parentheses, and a count (optionally added) is also a formula.

For example, "(H2O2)" and "(H2O2)3" are formulas.
Return the count of all elements as a string in the following form: the first name (in sorted order), followed by its count (if that count is more than 1), followed by the second name (in sorted order), followed by its count (if that count is more than 1), and so on.

The test cases are generated so that all the values in the output fit in a 32-bit integer.
*/

func countOfAtoms(formula string) string {
	if len(formula) == 0 {
		return ""
	}

	prev := make(map[string]int)
	ans := make(map[string]int)
	stack := make([]map[string]int, 0)

	for i := 0; i < len(formula); {
		if formula[i] >= 'A' && formula[i] <= 'Z' {
			element, l := nextElement(formula[i:])
			prev = make(map[string]int)
			prev[element] = 1
			ans = combine(prev, ans, 1)
			i += l
		} else if formula[i] >= '0' && formula[i] <= '9' {
			num, l := nextNum(formula[i:])
			ans = combine(prev, ans, num-1)
			i += l
		} else if formula[i] == '(' {
			stack = append(stack, ans)
			ans = make(map[string]int)
			i++
		} else if formula[i] == ')' {
			preAns := stack[len(stack)-1]
			prev = ans
			stack = stack[:len(stack)-1]
			ans = combine(ans, preAns, 1)
			i++
		}
	}

	atoms := make([]Atomic, len(ans))
	for k, v := range ans {
		atoms = append(atoms, Atomic{k, v})
	}

	sort.Slice(atoms, func(i, j int) bool {
		return atoms[i].ele < atoms[j].ele
	})

	res := ""
	for _, v := range atoms {
		res += v.ele
		if v.count > 1 {
			res += strconv.Itoa(v.count)
		}
	}
	return res
}

func nextElement(s string) (string, int) {
	for i := 1; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return s[:i], i
		}
	}
	return s, len(s)
}

func nextNum(s string) (int, int) {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num *= 10
			num += int(s[i] - '0')
		} else {
			return num, i
		}
	}
	return num, len(s)
}

func combine(pre, ans map[string]int, count int) map[string]int {
	for k, v := range pre {
		if value, ok := ans[k]; ok {
			ans[k] = value + v*count
		} else {
			ans[k] = v * count
		}
	}
	return ans
}

type Atomic struct {
	ele   string
	count int
}
