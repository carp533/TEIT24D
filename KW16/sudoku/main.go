package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var field = [9][9]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input string) [9][9]int {
	board := [9][9]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			i1, _ := strconv.Atoi(scanner.Text())
			board[row][col] = i1
		}
	}
	return board
}
func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	return false
}

func main() {
	field = parseInput("53..7....6..195....98....6.8...6...34..8.3..17...2...6.6....28....419..5....8..79")
	printBoard(field)
	if solve(0, 0) {
		fmt.Println("Found solution")
		printBoard(field)
	} else {
		fmt.Println("No solution")
	}
}
