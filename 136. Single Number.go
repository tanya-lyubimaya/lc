package lc

import "fmt"

func singleNumber(nums []int) int {
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
