package validsudoku

import (
	"strconv"
)

const empty byte = '.'

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	squares := [9][9]bool{}

	var (
		val       int64
		err       error
		squareIdx int
	)
	for i := range board {
		for j := range board[i] {
			elem := board[i][j]
			if elem == empty {
				continue
			}
			val, err = strconv.ParseInt(string(elem), 10, 8)
			if err != nil {
				panic(err)
			}
			val--

			squareIdx = (i/3)*3 + j/3

			if rows[i][val] || cols[j][val] || squares[squareIdx][val] {
				return false
			}

			rows[i][val] = true
			cols[j][val] = true
			squares[squareIdx][val] = true
		}
	}

	return true
}
