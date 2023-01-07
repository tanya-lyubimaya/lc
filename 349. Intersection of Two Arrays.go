package lc

func intersectionNaive(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 && len(nums2) == 0 && nums1[0] == nums2[0] {
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

// TODO: turn slices to sets by removing duplicates and use hash map to track intersections
// TODO: sort slices in advance and use pointers (try memory O(1))
