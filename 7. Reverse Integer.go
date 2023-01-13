package lc

import (
	"math"
	"strconv"
)

func reverse(x int) int {
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
