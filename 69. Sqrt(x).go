package lc

import "math"

func mySqrt(x int) int {
	i := 1
	for i*i <= x {
		i++
	}
	return i - 1
}

func mySqrtNewtonMethod(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	xPrev, xCurr := float64(x), float64(x/2)
	precision := 0.1
	for math.Abs(xPrev-xCurr) > precision {
		xPrev = xCurr
		xCurr = 0.5 * (xCurr + float64(x)/xCurr)
	}
	return int(xCurr)
}

// TODO: Binary search
