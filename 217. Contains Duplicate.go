package lc

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		if counter, _ := numsMap[v]; counter >= 1 {
			return true
		}
		numsMap[v]++
	}
	return false
}
