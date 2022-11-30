package DFS

import "fmt"

/*
Given an m x n integers matrix, return the length of the longest increasing path in matrix.

From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
*/

var dirs = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	path := 1
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	visited := make([][]int, len(matrix))
	for k := 0; k < len(matrix); k++ {
		visited[k] = make([]int, len(matrix[k]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result := dfs(i, j, matrix, dp, visited)
			path = max(path, result)
			dp[i][j] = result
		}
	}

	fmt.Println(dp)

	return path
}

func dfs(i, j int, matrix, dp, visited [][]int) int {
	if dp[i][j] > 0 {
		return dp[i][j]
	}

	visited[i][j] = 1
	path := 0
	for _, dir := range dirs {
		nextI := dir[0] + i
		nextJ := dir[1] + j
		if nextI >= 0 && nextI < len(matrix) && nextJ >= 0 && nextJ < len(matrix[0]) && visited[nextI][nextJ] == 0 && matrix[i][j] > matrix[nextI][nextJ] {
			path = max(path, dfs(nextI, nextJ, matrix, dp, visited))
		}
	}
	visited[i][j] = 0

	return path + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
