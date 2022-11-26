package sort

import "sort"

/*Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.

Return any possible rearrangement of s or return "" if not possible.



Example 1:

Input: s = "aab"
Output: "aba"
Example 2:

Input: s = "aaab"
Output: ""
*/

type bucket struct {
	c   byte
	num int
}

type bucketSlice []bucket

func (bs bucketSlice) Len() int           { return len(bs) }
func (bs bucketSlice) Less(i, j int) bool { return bs[i].num > bs[j].num }
func (bs bucketSlice) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }

func reorganizeString(S string) string {
	n := len(S)
	bs := make(bucketSlice, 26)

	for _, r := range S {
		b := r - 'a'
		if bs[b].num == 0 {
			bs[b].c = byte(r)
		}
		bs[b].num++
	}

	sort.Sort(bs)
	if bs[0].num > (n+1)/2 {
		return ""
	}
	res := make([]byte, n)
	j := 0
	for _, bucket := range bs {
		for i := 0; i < bucket.num; i, j = i+1, j+2 {
			if j >= n {
				j = 1
			}
			res[j] = bucket.c
		}
	}
	return string(res)
}
