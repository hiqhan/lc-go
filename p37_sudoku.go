package main

import "fmt"

const n int = 9
const e byte = '.'

func solveSudoku(board [][]byte) bool {
	row, col, ok := findEmptyCell(board)
	if !ok {
		return true
	}

	for i := 1; i <= 9; i++ {
		if isSafe(board, row, col, byte(48+i)) {
			board[row][col] = byte(48 + i)
			if solveSudoku(board) {
				return true
			}
			board[row][col] = e
		}
	}
	return false
}

func findEmptyCell(board [][]byte) (row, col int, ok bool) {
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == e {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

func isSafe(board [][]byte, row, col int, num byte) bool {
	return board[row][col] == e &&
		isRowSafe(board, row, num) &&
		isColSafe(board, col, num) &&
		isBoxSafe(board, row, col, num)
}

func isRowSafe(board [][]byte, row int, num byte) bool {
	for i := 0; i < n; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

func isColSafe(board [][]byte, col int, num byte) bool {
	for i := 0; i < n; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

func isBoxSafe(board [][]byte, row, col int, num byte) bool {
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func printBoard(b [][]byte) {
	for _, l := range b {
		fmt.Printf("%c\n", l)
	}
}

func main() {
	board := [][]byte{
		[]byte(string("53..7....")),
		[]byte(string("6..195...")),
		[]byte(string(".98....6.")),
		[]byte(string("8...6...3")),
		[]byte(string("4..8.3..1")),
		[]byte(string("7...2...6")),
		[]byte(string(".6....28.")),
		[]byte(string("...419..5")),
		[]byte(string("....8..79")),
	}
	if solveSudoku(board) {
		printBoard(board)
	}
}
