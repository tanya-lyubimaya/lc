package lc

import "strconv"

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

func isPowerOfThreeString(n int) bool {
	s := strconv.FormatInt(int64(n), 3)
	if s[0] != '1' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			return false
		}
	}
	return true
}

// TODO: Approach 3: Mathematics
// TODO: Solve with recursion
// TODO: Approach 4: Integer Limitations
