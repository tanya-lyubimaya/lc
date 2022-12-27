package lc

func findMaxConsecutiveOnesFirstAttempt(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var cnt int
	prev, maxCnt := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			if prev == 1 {
				if cnt < 2 {
					cnt = 2
				} else {
					cnt++
				}
			} else {
				cnt++
			}
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else {
			cnt = 0
		}
		prev = nums[i]
	}
	return maxCnt
}

func findMaxConsecutiveOnesBetter(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxCnt, cnt := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else {
			cnt = 0
		}
	}
	return maxCnt
}

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxCnt, cnt := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
			if cnt > maxCnt {
				maxCnt = cnt
			}
		} else {
			cnt = 0
		}
	}
	return maxCnt
}
