package lc

import "strconv"

func isPalindromeItoa(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var reverted int
	for x > reverted {
		reverted *= 10
		reverted += x % 10
		x /= 10
	}
	return x == reverted || x == reverted/10
}
