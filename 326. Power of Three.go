package lc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isPowerOfThreeNaive(n int) bool {
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

func isPowerOfThreeMathematics(n int) bool {
	val := math.Log10(float64(n)) / math.Log10(3)
	s := fmt.Sprintf("%v", val)
	split := strings.Split(s, ".")
	intPart, _ := strconv.ParseInt(split[0], 10, 0)
	return int(math.Pow(3, float64(intPart))) == n
}

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

// TODO: Solve with recursion
