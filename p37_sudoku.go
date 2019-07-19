package leetcode

import "fmt"

const N int = 9
const Empty byte = '.'

func solveSudoku(board [][]byte) bool {
	row, col, ok := findEmptyCell(board)
	if !ok {
		for _, l := range board {
			fmt.Println(l)
		}
		return true
	}

	for i := 1; i <= 9; i++ {
		if isSafe(board, row, col, byte(i)) {
			board[row][col] = byte(i)
			if solveSudoku(board) {
				return true
			}
			board[row][col] = Empty
		}
	}
	return false
}

func findEmptyCell(board [][]byte) (row, col int, ok bool) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == Empty {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

func isSafe(board [][]byte, row, col int, num byte) bool {
	return isRowSafe(board, row, num) && isColSafe(board, col, num) && isBoxSafe(board, row, col, num)
}

func isRowSafe(board [][]byte, row int, num byte) bool {
	for i := 0; i < N; i++ {
		if board[row][i] == num {
			return false
		}
	}
	return true
}

func isColSafe(board [][]byte, col int, num byte) bool {
	for i := 0; i < N; i++ {
		if board[i][col] == num {
			return false
		}
	}
	return true
}

func isBoxSafe(board [][]byte, row, col int, num byte) bool {
	startRow, startCol := row/3, col/3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}
