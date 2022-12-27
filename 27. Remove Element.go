package lc

func removeElement(nums []int, val int) int {
	k := len(nums)
	var i int
	for i < len(nums) {
		if nums[i] == val && k > i {
			for j := i; j < k-1; j++ {
				nums[j] = nums[j+1]
			}
			k--
		} else {
			i++
		}
	}
	return k
}
