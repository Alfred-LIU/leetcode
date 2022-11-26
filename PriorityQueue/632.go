package PriorityQueue

import "container/heap"

/*
You have k lists of sorted integers in non-decreasing order. Find the smallest range that includes at least one number from each of the k lists.

We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.



Example 1:

Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
Output: [20,24]
Explanation:
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].
*/

type min632Heap struct {
	nums      [][]int
	positions [][]int
}

func (h min632Heap) Len() int {
	return len(h.positions)
}
func (h min632Heap) Less(i, j int) bool {
	return h.nums[h.positions[i][0]][h.positions[i][1]] < h.nums[h.positions[j][0]][h.positions[j][1]]
}
func (h min632Heap) Swap(i, j int) {
	h.positions[i], h.positions[j] = h.positions[j], h.positions[i]
}
func (h *min632Heap) Push(i interface{}) {
	(*h).positions = append((*h).positions, i.([]int))
}
func (h *min632Heap) Pop() interface{} {
	l := (*h).positions[len((*h).positions)-1]
	(*h).positions = (*h).positions[:len((*h).positions)-1]
	return l
}
func smallestRange(nums [][]int) []int {
	h := min632Heap{
		nums:      nums,
		positions: [][]int{},
	}
	mn, mx := nums[0][0], nums[0][0]
	for i := 0; i < len(nums); i++ {
		heap.Push(&h, []int{i, 0})
		mn = min(mn, nums[i][0])
		mx = max(mx, nums[i][0])
	}

	res := []int{mn, mx}
	for true {
		position := heap.Pop(&h).([]int)
		i, j := position[0], position[1]
		j++
		if j == len(nums[i]) {
			return res
		}

		heap.Push(&h, []int{i, j})
		mn, mx = nums[h.positions[0][0]][h.positions[0][1]], max632(mx, nums[i][j])
		if mx-mn < res[1]-res[0] || (mn < res[0] && mx-mn == res[1]-res[0]) {
			res = []int{mn, mx}
		}
	}
	return []int{}
}

func max632(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
