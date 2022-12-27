package lc

func checkIfExist(arr []int) bool {
	elements := make(map[int]int, len(arr))
	for i, v := range arr {
		elements[v] = i
		if _, ok := elements[v*2]; ok {
			return true
		}
		if _, ok := elements[v/2]; ok && v%2 == 0 {
			return true
		}
	}
	return false
}
