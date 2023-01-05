package lc

import (
	"fmt"
	"sort"
)

func singleNumberHashMap(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}
	fmt.Println(numsMap)
	for k, v := range numsMap {
		fmt.Println(v)
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberSort(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i++ {
		if i == 1 && nums[i] != nums[i-1] {
			return nums[i-1]
		}
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
