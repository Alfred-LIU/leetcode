package Math

func maxChunksToSorted(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxElement := arr[0]
	maxArray := make([]int, 0)
	maxArray = append(maxArray, arr[0])

	for i := 1; i < len(arr); i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
		maxArray = append(maxArray, maxElement)
	}

	res := 0
	for i := 0; i < len(maxArray); i++ {
		if maxArray[i] == i {
			res += 1
		}
	}

	return res
}
