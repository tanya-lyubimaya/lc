package lc

func removeDuplicates(nums []int) int {
	i := 1
	k := len(nums)
	for i < len(nums) && k > i {
		if nums[i] == nums[i-1] {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			k--
		} else {
			i++
		}
	}
	return k
}
