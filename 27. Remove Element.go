package lc

func removeElementNaive(nums []int, val int) int {
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

func removeElement(nums []int, val int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return j
}

func removeRareElement(nums []int, val int) int {
	k := len(nums)
	var i int
	for i < k {
		if nums[i] == val {
			nums[i] = nums[k-1]
			k--
		} else {
			i++
		}
	}
	return k
}
