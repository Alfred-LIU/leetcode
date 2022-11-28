package bruteForce

/*
This question is about implementing a basic elimination algorithm for Candy Crush.

Given an m x n integer array board representing the grid of candy where board[i][j] represents the type of candy. A value of board[i][j] == 0 represents that the cell is empty.

The given board represents the state of the game following the player's move. Now, you need to restore the board to a stable state by crushing candies according to the following rules:

If three or more candies of the same type are adjacent vertically or horizontally, crush them all at the same time - these positions become empty.
After crushing all candies simultaneously, if an empty space on the board has candies on top of itself, then these candies will drop until they hit a candy or bottom at the same time. No new candies will drop outside the top boundary.
After the above steps, there may exist more candies that can be crushed. If so, you need to repeat the above steps.
If there does not exist more candies that can be crushed (i.e., the board is stable), then return the current board.
You need to perform the above rules until the board becomes stable, then return the stable board.
*/

func candyCrush(board [][]int) [][]int {
	n := len(board)
	m := len(board[0])

	found := true

	for found {
		found = false
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				v := abs(board[i][j])

				if v == 0 {
					continue
				}

				if j < m-2 && abs(board[i][j+1]) == v && abs(board[i][j+2]) == v {
					found = true
					itr := j
					for itr < m {
						if abs(board[i][itr]) != v {
							break
						}
						board[i][itr] = -v
						itr++
					}
				}

				if i < n-2 && abs(board[i+1][j]) == v && abs(board[i+2][j]) == v {
					found = true
					itr := i
					for itr < n {
						if abs(board[itr][j]) != v {
							break
						}
						board[itr][j] = -v
						itr++
					}
				}

			}
		}

		if found {
			for j := 0; j < m; j++ {
				itr := n - 1
				for i := n - 1; i >= 0; i-- {
					if board[i][j] > 0 {
						board[itr][j] = board[i][j]
						itr--
					}
				}
				for i := itr; i >= 0; i-- {
					board[i][j] = 0
				}
			}

		}
	}

	return board
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
