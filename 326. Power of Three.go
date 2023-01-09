package lc

func isPowerOfThree(n int) bool {
	res := 1
	for res < n {
		res *= 3
	}
	if res == n {
		return true
	}
	return false
}

// TODO: Solve with recursion
// TODO: Solve without loops or recursion
