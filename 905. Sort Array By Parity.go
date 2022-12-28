package lc

func sortArrayByParity(nums []int) []int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
