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

func singleNumberSum(nums []int) int {
	numsMap := make(map[int]int)
	var numsSum, numsMapSum int
	for _, v := range nums {
		numsSum += v
		if _, ok := numsMap[v]; !ok {
			numsMap[v]++
		}
	}
	for k, _ := range numsMap {
		numsMapSum += k
	}
	return numsMapSum*2 - numsSum
}

func singleNumber(nums []int) int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	return xor
}
