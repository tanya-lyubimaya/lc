package lc

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	characters := make(map[rune]int)
	for _, c := range s {
		characters[c]++
	}
	for i, c := range s {
		for k, v := range characters {
			if v == 1 && c == k {
				return i
			}
		}
	}
	return -1
}
