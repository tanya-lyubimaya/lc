package lc

import "strconv"

func findNumbersString(nums []int) int {
	var cnt int
	for _, v := range nums {
		if len(strconv.Itoa(v))%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func findNumbers(nums []int) int {
	var cnt int
	for _, v := range nums {
		var nDigits int
		temp := float32(v)
		for temp >= 1 {
			temp /= 10
			nDigits++
		}
		if nDigits%2 == 0 {
			cnt++
		}
	}
	return cnt
}
