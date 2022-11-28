package DFS

/*
Given an array of non-negative integers arr, you are initially positioned at start index of the array. When you are at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach to any index with value 0.

Notice that you can not jump outside of the array at any time.
*/
func canReach(arr []int, start int) bool {
	if start >= len(arr) || start < 0 {
		return false
	}

	val := arr[start]

	if val == 0 {
		return true
	}

	if val < 0 {
		return false
	}

	arr[start] *= -1

	return canReach(arr, start+val) || canReach(arr, start-val)
}
