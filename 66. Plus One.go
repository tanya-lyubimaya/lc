package lc

func plusOne(digits []int) []int {
	overflow := 1
	for i := len(digits) - 1; i >= 0 && overflow != 0; i-- {
		if digits[i]+overflow > 9 {
			overflow = digits[i] + overflow - 9
			digits[i] = 0
		} else {
			digits[i] += overflow
			overflow = 0
		}
	}
	if overflow > 0 {
		digits[0] = 0
		digits = append([]int{overflow}, digits...)
	}
	return digits
}
