package lc

import "strconv"

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

func isValidSudokuBitSetting(board [][]byte) bool {
	for _, row := range board {
		var n int16
		for _, v := range row {
			if v != '.' {
				bitPos, _ := strconv.Atoi(string(v))
				if !hasBit(n, bitPos) {
					n = setBit(n, bitPos)
				} else {
					return false
				}
			}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		var n int16
		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' {
				bitPos, _ := strconv.Atoi(string(board[j][i]))
				if !hasBit(n, bitPos) {
					n = setBit(n, bitPos)
				} else {
					return false
				}
			}
		}
	}
	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			var n int16
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] != '.' {
						bitPos, _ := strconv.Atoi(string(board[x][y]))
						if !hasBit(n, bitPos) {
							n = setBit(n, bitPos)
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

func setBit(n int16, pos int) int16 {
	n |= (1 << pos)
	return n
}

/*
	func clearBit(n int16, pos int) int16 {
	    mask := ^(1 << pos)
	    n &= mask
	    return n
	}
*/
func hasBit(n int16, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}
