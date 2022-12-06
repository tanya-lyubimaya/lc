package lc

func runningSumClassic(nums []int) []int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}

	return sums
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	
	return nums
}
