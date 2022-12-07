package lc

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, v := range accounts {
		sum := 0
		for _, v2 := range v {
			sum += v2
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}
