package lc

import "sort"

func sortArrayByParity(nums []int) []int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func sortArrayByParitySliceStable(nums []int) []int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i]%2 == 0 && nums[j]%2 == 1
	})
	return nums
}
