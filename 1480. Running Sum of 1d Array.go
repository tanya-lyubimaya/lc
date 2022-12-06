package lc

func runningSumNaive(nums []int) []int {
	sums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := 0; j <= i; j++ {
			temp += nums[j]
		}
		sums[i] = temp
	}
	return sums
}

func runningSum(nums []int) []int {
	sums := make([]int, len(nums))
	var temp int
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		sums[i] = temp
	}
	return sums
}
