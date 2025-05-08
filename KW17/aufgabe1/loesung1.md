
// isSafe checks if a queen can be placed at board[row][col]
// It checks if there's any queen in the same column,
// upper diagonal, or lower diagonal
func isSafe(board [N][N]bool, row, col int) bool {
	// Check this column on upper rows
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}

	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < N; i, j = i-1, j+1 {
		if board[i][j] {
			return false
		}
	}

	return true
}

// solveNQueensUtil is a recursive function to solve the N Queens problem
func solveNQueensRec(board *[N][N]bool, row int, solutions *[][N][N]bool) bool {
	// Base case: If all queens are placed
	if row == N {
		// Add this solution to our list of solutions
		*solutions = append(*solutions, *board)
		return true
	}

	// Consider this row and try placing the queen in all columns one by one
	found := false
	for col := 0; col < N; col++ {
		// Check if the queen can be placed at board[row][col]
		if isSafe(*board, row, col) {
			// Place the queen
			board[row][col] = true

			// Recursively try to place the rest of the queens
			if solveNQueensRec(board, row+1, solutions) {
				found = true
			}

			// If placing the queen doesn't lead to a solution, backtrack
			board[row][col] = false
		}
	}

	return found
}

