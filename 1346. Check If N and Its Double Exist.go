package lc

func checkIfExist(arr []int) bool {
	type void struct{}
	var member void
	elements := make(map[int]void)
	for _, v := range arr {
		if _, ok := elements[v*2]; ok {
			return true
		}
		if _, ok := elements[v/2]; ok && v%2 == 0 {
			return true
		}
		elements[v] = member
	}
	return false
}
