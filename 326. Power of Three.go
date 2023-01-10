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

func isPowerOfThreeLoop(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// TODO: Approach 2: Base Conversion
// TODO: Approach 3: Mathematics
// TODO: Solve with recursion
// TODO: Approach 4: Integer Limitations
