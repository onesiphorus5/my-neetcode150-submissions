func isValidSudoku(board [][]byte) bool {
	table := make([]int, 10)
	var v byte
	for i := 0; i<9; i++ {
		// walk row
		clear(table)
		for j := 0; j<9; j++ {
			v = board[i][j]
			if v != '.' {
				if table[v-'0'] >= 1 {
					return false
				}
				table[v-'0'] += 1
			}
		}

		// walk colum
		clear(table)
		for j := 0; j<9; j++ {
			v = board[j][i]
			if v != '.' {
				if table[v-'0'] >= 1 {
					return false
				}
				table[v-'0'] += 1
			}
		}
	}

	// walk the 9 boxes
	for i := 0; i < 9; i = i+3 {
		for j := 0; j < 9; j = j+3 {
			// work the {i,j} box
			clear(table)
			for zi := i; zi < (i+3); zi++ {
				for zj := j; zj < (j+3); zj++ {
					v = board[zi][zj]
					if v != '.' {
				    	if table[v-'0'] >= 1 {
							return false
						}
						table[v-'0'] += 1
					}
				}
			}
		}
	}
	return true
}
