package lc

import "strconv"

func hammingWeight(num uint32) int {
	var cnt int
	s := strconv.FormatInt(int64(num), 2)
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}
	return cnt
}
