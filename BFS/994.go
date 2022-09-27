package BFS

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.
*/

var dir = [][]int{
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
	[]int{-1, 0},
}

func orangesRotting(grid [][]int) int {

	move := 0
	target := 0
	pos := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				pos = append(pos, []int{i, j})
			} else if grid[i][j] == 1 {
				target++
			}

		}
	}

	if target == 0 {
		return 0
	}

	for len(pos) > 0 {
		if target == 0 {
			return move
		}
		newPos := [][]int{}

		for _, v := range pos {
			for _, d := range dir {
				ni, nj := v[0]+d[0], v[1]+d[1]
				if ni < 0 || nj < 0 || ni >= len(grid) || nj >= len(grid[0]) || grid[ni][nj] != 1 {
					continue
				}
				grid[ni][nj] = 2
				newPos = append(newPos, []int{ni, nj})
				target--
			}
		}

		pos = newPos
		move++
	}

	return -1
}
