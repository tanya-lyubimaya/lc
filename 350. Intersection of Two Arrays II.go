package lc

import "sort"

func intersectHashMap(nums1 []int, nums2 []int) []int {
	if len(nums1) == 1 && len(nums2) == 1 && nums1[0] == nums2[0] {
		return []int{nums1[0]}
	}
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)
	var result []int
	for _, v := range nums1 {
		nums1Map[v]++
	}
	for _, v := range nums2 {
		nums2Map[v]++
	}
	for k, v := range nums1Map {
		if v2, ok := nums2Map[k]; ok {
			if v < v2 {
				for i := 0; i < v; i++ {
					result = append(result, k)
				}
			} else {
				for i := 0; i < v2; i++ {
					result = append(result, k)
				}
			}
		}
	}
	return result
}

func intersect(nums1 []int, nums2 []int) []int {
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
	}
	return result
}
