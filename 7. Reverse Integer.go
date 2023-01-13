package lc

import (
	"math"
	"strconv"
)

func reverseStringConversion(x int) int {
	str := strconv.Itoa(x)
	var resStr []byte
	if str[0] == '-' {
		resStr = append(resStr, str[0])
	}
	for i := len(str) - 1; i >= 0 && str[i] != '-'; i-- {
		resStr = append(resStr, str[i])
	}
	res, _ := strconv.Atoi(string(resStr))
	if res > math.MaxInt32 || res < -1*math.MaxInt32 {
		return 0
	}
	return res
}

func reverse(x int) int {
	var res int
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && pop < -8) {
			return 0
		}
		res *= 10
		res += pop
	}
	return res
}
