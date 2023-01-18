package lc

import "sort"

func intersectionLoops(nums1 []int, nums2 []int) []int {
	if len(nums1) == 1 && len(nums2) == 1 && nums1[0] == nums2[0] {
		return []int{nums1[0]}
	}
	var result []int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				var hasDuplicate bool
				for _, v := range result {
					if v == nums1[i] {
						hasDuplicate = true
					}
				}
				if !hasDuplicate {
					result = append(result, nums1[i])
				}
			}
		}
	}
	return result
}

func intersectionHashMap(nums1 []int, nums2 []int) []int {
	if len(nums1) == 1 && len(nums2) == 1 && nums1[0] == nums2[0] {
		return []int{nums1[0]}
	}
	numsMap := make(map[int]int)
	k1, k2 := len(nums1), len(nums2)
	for i := 1; i < k1; i++ {
		if nums1[i] == nums1[i-1] {
			for j := i; j < k1-1; j++ {
				nums1[j] = nums1[j+1]
			}
			k1--
		}
	}
	for i := 1; i < k2; i++ {
		if nums2[i] == nums2[i-1] {
			for j := i; j < k2-1; j++ {
				nums2[j] = nums2[j+1]
			}
			k2--
		}
	}
	for i := 0; i < k1; i++ {
		if _, ok := numsMap[nums1[i]]; !ok {
			numsMap[nums1[i]]++
		}
	}
	for i := 0; i < k2; i++ {
		if v, _ := numsMap[nums2[i]]; v == 1 {
			numsMap[nums2[i]]++
		}
	}
	nums1 = nil
	for k, v := range numsMap {
		if v == 2 {
			nums1 = append(nums1, k)
		}
	}
	return nums1
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 1 && len(nums2) == 1 && nums1[0] == nums2[0] {
		return []int{nums1[0]}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j int
	var result []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
		for i > 0 && i < len(nums1) && nums1[i] == nums1[i-1] {
			i++
		}
		for j > 0 && j < len(nums2) && nums2[j] == nums2[j-1] {
			j++
		}
	}
	return result
}
