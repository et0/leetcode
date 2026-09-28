package main // https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	squads := [3][3][9]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			key := board[i][j] - '1'

			if rows[i][key] {
				return false
			}
			rows[i][key] = true

			if cols[j][key] {
				return false
			}
			cols[j][key] = true

			if squads[i/3][j/3][key] {
				return false
			}
			squads[i/3][j/3][key] = true
		}
	}

	return true
}
