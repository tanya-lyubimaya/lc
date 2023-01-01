package lc

func findDisappearedNumbers(nums []int) []int {
	numsMap := make(map[int]int, len(nums))
	var cnt int
	for _, v := range nums {
		numsMap[v]++
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := numsMap[i]; !ok {
			nums[cnt] = i
			cnt++
		}
	}
	return nums[:cnt]
}
