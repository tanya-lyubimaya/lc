package lc

func rotate(nums []int, k int) {
	for i := 1; i <= k; i++ {
		lastNum := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = lastNum
	}
}
