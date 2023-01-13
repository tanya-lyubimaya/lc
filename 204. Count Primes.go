package lc

import "math"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	compositeArr := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if compositeArr[i] {
			continue
		}
		count++
		if float64(i) > math.Sqrt(float64(n)) {
			continue
		}
		for j := i * i; j < n; j += i * 2 {
			compositeArr[j] = true
		}
	}
	return count
}
