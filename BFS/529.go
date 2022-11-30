package BFS

/*
Let's play the minesweeper game (Wikipedia, online game)!

You are given an m x n char matrix board representing the game board where:

'M' represents an unrevealed mine,
'E' represents an unrevealed empty square,
'B' represents a revealed blank square that has no adjacent mines (i.e., above, below, left, right, and all 4 diagonals),
digit ('1' to '8') represents how many mines are adjacent to this revealed square, and
'X' represents a revealed mine.
You are also given an integer array click where click = [clickr, clickc] represents the next click position among all the unrevealed squares ('M' or 'E').

Return the board after revealing this position according to the following rules:

If a mine 'M' is revealed, then the game is over. You should change it to 'X'.
If an empty square 'E' with no adjacent mines is revealed, then change it to a revealed blank 'B' and all of its adjacent unrevealed squares should be revealed recursively.
If an empty square 'E' with at least one adjacent mine is revealed, then change it to a digit ('1' to '8') representing the number of adjacent mines.
Return the board when no more squares will be revealed.
*/

func updateBoard(board [][]byte, click []int) [][]byte {
	var queue [][2]int
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	if board[click[0]][click[1]] == 'E' {
		count := countSurroundingMines(board, click[0], click[1])
		if count == 0 {
			board[click[0]][click[1]] = 'B'
			queue = append(queue, [2]int{click[0], click[1]})
		} else {
			board[click[0]][click[1]] = byte(count + '0')
		}
	}
	var curr [2]int
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		row, col := curr[0], curr[1]
		for _, dir := range directions {
			x, y := row+dir[0], col+dir[1]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) && board[x][y] == 'E' {
				count := countSurroundingMines(board, x, y)
				if count == 0 {
					board[x][y] = 'B'
					// Add to the queue to recursively reveal surroundings
					queue = append(queue, [2]int{x, y})
				} else {
					// Found a neighboring mine so don't reveal surroundings
					board[x][y] = byte(count + '0')
				}
			}
		}
	}
	return board
}

var directions = [8][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func countSurroundingMines(board [][]byte, row, col int) int {
	var res int
	for _, dir := range directions {
		x, y := row+dir[0], col+dir[1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) && board[x][y] == 'M' {
			res++
		}
	}
	return res
}
