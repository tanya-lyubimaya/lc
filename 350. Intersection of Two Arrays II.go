package lc

func intersect(nums1 []int, nums2 []int) []int {
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
