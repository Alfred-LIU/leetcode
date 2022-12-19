package Math

func maxChunksToSorted1(arr []int) int {
	l := len(arr)
	maxLeft := make([]int, l)
	minRight := make([]int, l)

	maxLeft[0] = arr[0]
	for i := 1; i < l; i++ {
		maxLeft[i] = max(maxLeft[i-1], arr[i])
	}

	minRight[l-1] = arr[l-1]
	for i := l - 2; i >= 0; i-- {
		minRight[i] = min(minRight[i+1], arr[i])
	}

	res := 0
	for i := 0; i < l-1; i++ {
		if maxLeft[i] <= minRight[i+1] {
			res++
		}
	}

	return res + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
