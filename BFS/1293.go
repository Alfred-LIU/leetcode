package BFS

/*
You are given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle). You can move up, down, left, or right from and to an empty cell in one step.

Return the minimum number of steps to walk from the upper left corner (0, 0) to the lower right corner (m - 1, n - 1) given that you can eliminate at most k obstacles. If it is not possible to find such walk return -1.
*/

// define moves in grid
var moves = [][]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

func shortestPath(grid [][]int, k int) int {
	stack := make([][]int, 0)
	// define visited path to avoid pushing stack repeatedly
	visited := make([][]int, len(grid))
	// use -1 for never visit
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	visited[0][0] = k
	stack = append(stack, []int{0, 0, k})
	level := 0
	for len(stack) != 0 {
		size := len(stack)
		for i := range stack {
			if stack[i][0] == len(grid)-1 && stack[i][1] == len(grid[0])-1 {
				return level
			}
			for _, move := range moves {
				nextx := stack[i][0] + move[0]
				nexty := stack[i][1] + move[1]
				// check if next move is valid
				if nextx >= 0 && nextx < len(grid) && nexty >= 0 && nexty < len(grid[0]) {
					curK := stack[i][2]
					if grid[nextx][nexty] == 1 {
						curK--
					}
					// chance left for eliminating obstacle
					if curK >= 0 {
						// use largest chance for visited and push stack
						if visited[nextx][nexty] == -1 || (visited[nextx][nexty] != -1 && curK > visited[nextx][nexty]) {
							visited[nextx][nexty] = curK
							stack = append(stack, []int{nextx, nexty, curK})
						}
					}
				}
			}
		}
		stack = stack[size:]
		level++
	}
	return -1
}