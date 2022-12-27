package lc

import (
	"math"
	"sort"
)

func sortedSquares(nums []int) []int {
	var squares []int
	for _, v := range nums {
		squares = append(squares, int(math.Pow(float64(v), 2)))
	}
	sort.Ints(squares)
	return squares
}
