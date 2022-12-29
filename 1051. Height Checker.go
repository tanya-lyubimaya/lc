package lc

import "sort"

func heightChecker(heights []int) int {
	var cnt int
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			cnt++
		}
	}
	return cnt
}

func heightCheckerFrequencies(heights []int) int {
	heightFreqs := make(map[int]int, len(heights))
	var res, curHeight int
	for _, v := range heights {
		heightFreqs[v]++
	}
	for i := 0; i < len(heights); i++ {
		for heightFreqs[curHeight] == 0 {
			curHeight++
		}
		if curHeight != heights[i] {
			res++
		}
		heightFreqs[curHeight]--
	}
	return res
}
