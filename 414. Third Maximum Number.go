package lc

import (
	"math"
	"sort"
)

func thirdMaxSorting(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	elemCounted := 1
	prevElem := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prevElem {
			elemCounted++
			prevElem = nums[i]
		}
		if elemCounted == 3 {
			return nums[i]
		}
	}
	return nums[0]
}

func thirdMaxPointers(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxs := [3]int{math.MinInt, math.MinInt, math.MinInt}

	for _, v := range nums {
		if maxs[0] == v || maxs[1] == v || maxs[2] == v {
			continue
		}
		if v > maxs[0] {
			maxs[2] = maxs[1]
			maxs[1] = maxs[0]
			maxs[0] = v
		} else if v > maxs[1] {
			maxs[2] = maxs[1]
			maxs[1] = v
		} else if v > maxs[2] {
			maxs[2] = v
		}
	}
	if maxs[2] == math.MinInt {
		return maxs[0]
	} else {
		return maxs[2]
	}
}

func thirdMaxPointersAndBooleans(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	firstMax, secondMax, thirdMax := [2]int{}, [2]int{}, [2]int{}

	for _, v := range nums {
		if (firstMax[0] == v && firstMax[1] == 1) || (secondMax[0] == v && secondMax[1] == 1) || (thirdMax[0] == v && thirdMax[1] == 1) {
			continue
		}
		if v > firstMax[0] || firstMax[1] == 0 {
			thirdMax[0], thirdMax[1] = secondMax[0], secondMax[1]
			secondMax[0], secondMax[1] = firstMax[0], firstMax[1]
			firstMax[0], firstMax[1] = v, 1
		} else if v > secondMax[0] || secondMax[1] == 0 {
			thirdMax[0], thirdMax[1] = secondMax[0], secondMax[1]
			secondMax[0], secondMax[1] = v, 1
		} else if v > thirdMax[0] || thirdMax[1] == 0 {
			thirdMax[0], thirdMax[1] = v, 1
		}
	}

	if thirdMax[1] == 0 {
		return firstMax[0]
	} else {
		return thirdMax[0]
	}
}

// TODO: Approach 2: Min Heap Data Structure
// TODO: Approach 3: Ordered Set
