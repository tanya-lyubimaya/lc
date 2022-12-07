package lc

func numberOfSteps(num int) int {
	var stepCount int
	for num != 0 {
		if num%2 == 0 {
			stepCount++
			num /= 2
		}
		if num%2 != 0 {
			stepCount++
			num -= 1
		}
	}
	return stepCount
}
