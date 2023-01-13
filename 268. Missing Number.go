package lc

import "sort"

func missingNumberSort(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

func missingNumberMap(nums []int) int {
	type void struct{}
	var member void
	m := make(map[int]void)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = member
		}
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return len(nums)
}

func missingNumber(nums []int) int {
	var sum int
	for i := 0; i <= len(nums); i++ {
		sum += i
	}
	for _, v := range nums {
		sum -= v
	}
	return sum
}
