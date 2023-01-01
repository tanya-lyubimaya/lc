package lc

import (
	"math"
	"sort"
)

func sortedSquaresFirstAttempt(nums []int) []int {
	var squares []int
	for _, v := range nums {
		squares = append(squares, int(math.Pow(float64(v), 2)))
	}
	sort.Ints(squares)
	return squares
}

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = int(math.Pow(float64(v), 2))
	}
	sort.Ints(nums)
	return nums
}
