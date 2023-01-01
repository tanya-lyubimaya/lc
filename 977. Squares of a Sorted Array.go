package lc

import (
	"math"
	"sort"
)

func sortedSquaresWithAdditionalArray(nums []int) []int {
	squares := make([]int, len(nums))
	for i, v := range nums {
		squares[i] = int(math.Pow(float64(v), 2))
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
