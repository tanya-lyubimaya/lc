package lc

import "sort"

func mergeTrivial(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			k--
			j--
		} else {
			nums1[k] = nums1[i]
			k--
			i--
		}
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
