package main

/*
Aufgabe1: Programmiere das 8 Damen Problem (mit Backtracking Algorithmus)
siehe: https://de.wikipedia.org/wiki/Damenproblem
Beispiel Lösung:
 ♕ .  .  .  .  .  .  .
 .  .  .  .  ♕ .  .  .
 .  .  .  .  .  .  .  ♕
 .  .  .  .  .  ♕ .  .
 .  .  ♕ .  .  .  .  .
 .  .  .  .  .  .  ♕ .
 .  ♕ .  .  .  .  .  .
 .  .  .  ♕ .  .  .  .

 Das unten stehende Coding kann als Hilfe genutzt werden.

*/

import (
	"fmt"
)

// The size of the chessboard (8x8 for 8 queens)
const N = 8
const white_queen = '\u2655'

func printSolution(board [N][N]bool) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] {
				fmt.Printf(" %c ", white_queen)
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// isSafe checks if a queen can be placed at board[row][col]
// It checks if there's any queen in the same column,
// upper diagonal, or lower diagonal
func isSafe(board [N][N]bool, row, col int) bool {
	// Check this column on upper rows
	//TODO

	// Check upper left diagonal
	//TODO

	// Check upper right diagonal
	//TODO

	return true
}

func solveNQueensRec(board *[N][N]bool, row int, solutions *[][N][N]bool) bool {
	if row == N {
		// Add this solution to our list of solutions
		*solutions = append(*solutions, *board)
		return true
	}

	found := false
	for col := 0; col < N; col++ {
		//TODO
	}

	return found
}

func solveNQueens() [][N][N]bool {
	var board [N][N]bool
	var solutions [][N][N]bool

	solveNQueensRec(&board, 0, &solutions)
	return solutions
}

func main() {
	solutions := solveNQueens()
	fmt.Printf("%d Lösungen für das %d-Damen Problem.\n\n", len(solutions), N)

	maxToShow := 3
	if len(solutions) < maxToShow {
		maxToShow = len(solutions)
	}

	for i := 0; i < maxToShow; i++ {
		fmt.Printf("Lösung %d:\n", i+1)
		printSolution(solutions[i])
	}
}
