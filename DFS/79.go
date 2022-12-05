package DFS

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.
*/

var dires79 = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res := helper79(i, j, 0, board, word)
			if res {
				return true
			}
		}
	}

	return false
}

func helper79(i, j, w int, board [][]byte, word string) bool {
	if w == len(word) {
		return true
	}
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}

	if board[i][j] != word[w] {
		return false
	}

	original := board[i][j]
	board[i][j] = '@'

	res := false
	for _, dir := range dirs {
		nextI := i + dir[0]
		nextJ := j + dir[1]
		res = res || helper79(nextI, nextJ, w+1, board, word)
	}
	if res {
		return true
	}

	board[i][j] = original

	return false
}
