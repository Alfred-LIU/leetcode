package array

//1424. Diagonal Traverse II: Given a 2D integer array nums, return all elements of nums in diagonal order as shown in the below images.
func findDiagonalOrder(nums [][]int) []int {
	arr := [][]int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j >= len(arr) {
				arr = append(arr, []int{})
			}

			arr[i+j] = append(arr[i+j][:0], append([]int{nums[i][j]}, arr[i+j][0:]...)...)
		}
	}

	for k := 0; k < len(arr); k++ {
		res = append(res, arr[k]...)
	}

	return res
}
