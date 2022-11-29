package rand

import "math/rand"

/*
You are given an integer n and an array of unique integers blacklist. Design an algorithm to pick a random integer in the range [0, n - 1] that is not in blacklist. Any integer that is in the mentioned range and not in blacklist should be equally likely to be returned.

Optimize your algorithm such that it minimizes the number of calls to the built-in random function of your language.

Implement the Solution class:

Solution(int n, int[] blacklist) Initializes the object with the integer n and the blacklisted integers blacklist.
int pick() Returns a random integer in the range [0, n - 1] and not in blacklist.
*/

type Solution struct {
	size    int
	replace map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	size := n - len(blacklist)
	replace := make(map[int]int)

	for _, v := range blacklist {
		replace[v] = -1
	}

	N := n
	for _, v := range blacklist {
		if v < size {
			for N > 0 {
				N--
				if _, ok := replace[N]; !ok {
					replace[v] = N
					break
				}
			}
		}
	}

	return Solution{
		size:    size,
		replace: replace,
	}
}

func (this *Solution) Pick() int {
	r := rand.Intn(this.size)
	if v, ok := this.replace[r]; ok {
		return v
	}
	return r
}
