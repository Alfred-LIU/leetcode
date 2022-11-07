package set

/*
You are given a non-negative integer array nums. In one operation, you must:

Choose a positive integer x such that x is less than or equal to the smallest non-zero element in nums.
Subtract x from every positive element in nums.
Return the minimum number of operations to make every element in nums equal to 0.
*/

func minimumOperations(nums []int) int {
	mapNums := make(map[int]struct{})
	for i := range nums {
		mapNums[nums[i]] = struct{}{}
	}
	if _, ok := mapNums[0]; ok {
		return len(mapNums) - 1
	}
	return len(mapNums)
}
