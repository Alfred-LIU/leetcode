package dp

/*
Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.
Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
Output: 3
Explanation: The repeated subarray with maximum length is [3,2,1].
*/

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	res := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				res = max(res, dp[i+1][j+1])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
