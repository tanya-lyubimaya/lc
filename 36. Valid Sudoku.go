package lc

func isValidSudoku(board [][]byte) bool {
	type void struct{}
	var member void
	for _, row := range board {
		rowMap := make(map[byte]void)
		for _, v := range row {
			if v != '.' {
				if _, ok := rowMap[v]; !ok {
					rowMap[v] = member
				} else {
					return false
				}
			}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		columnMap := make(map[byte]void)
		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' {
				if _, ok := columnMap[board[j][i]]; !ok {
					columnMap[board[j][i]] = member
				} else {
					return false
				}
			}
		}
	}
	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			gridMap := make(map[byte]void)
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] != '.' {
						if _, ok := gridMap[board[x][y]]; !ok {
							gridMap[board[x][y]] = member
						} else {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
