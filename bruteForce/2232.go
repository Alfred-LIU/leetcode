package bruteForce

import (
	"fmt"
	"strconv"
)

/*
You are given a 0-indexed string expression of the form "<num1>+<num2>" where <num1> and <num2> represent positive integers.

Add a pair of parentheses to expression such that after the addition of parentheses, expression is a valid mathematical expression and evaluates to the smallest possible value. The left parenthesis must be added to the left of '+' and the right parenthesis must be added to the right of '+'.

Return expression after adding a pair of parentheses such that expression evaluates to the smallest possible value. If there are multiple answers that yield the same result, return any of them.

The input has been generated such that the original value of expression, and the value of expression after adding any pair of parentheses that meets the requirements fits within a signed 32-bit integer.
*/

func minimizeResult(expression string) string {
	m := map[int]string{}
	s := strings.Split(expression, "+")
	n1 := s[0]
	n2 := s[1]

	// get the normal sum
	res, _ := strconv.Atoi(s[0])
	res2, _ := strconv.Atoi(s[1])
	res += res2

	//map it with the expression
	m[res] = fmt.Sprintf("(%s)", expression)

	for i := 0; i < len(n1); i++ {
		i1, err := strconv.Atoi(n1[:i])
		if err != nil {
			i1 = 1
		}
		i2, _ := strconv.Atoi(n1[i:])
		//fmt.Println("i",i1,i2)
		for j := 1; j <= len(n2); j++ {
			j1, _ := strconv.Atoi(n2[:j])
			j2, err := strconv.Atoi(n2[j:])
			if err != nil {
				j2 = 1
			}
			//fmt.Println("j",j1,j2)

			// calculate
			curr := i1 * (i2 + j1) * j2
			m[curr] = fmt.Sprintf("%[1]s(%[2]s+%[3]s)%[4]s", n1[:i], n1[i:], n2[:j], n2[j:])
			if (curr < res) {
				res = curr
			}

		}
	}

	return m[res]
}