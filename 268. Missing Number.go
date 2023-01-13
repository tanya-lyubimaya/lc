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

func missingNumberSum(nums []int) int {
	var sum int
	for i := 0; i <= len(nums); i++ {
		sum += i
	}
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func missingNumberBinarySearch(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	mid := (left + right) / 2
	for left < right {
		mid = (left + right) / 2
		if nums[mid] > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
