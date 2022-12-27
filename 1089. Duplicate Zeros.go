package lc

func duplicateZerosNestedAppends(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(arr[:i], append([]int{0}, arr[i:len(arr)-1]...)...)
			i++
		}
	}
}

func duplicateZerosAppend(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}

func duplicateZeros(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] == 0 {
			for j := l - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
}

func duplicateZerosCopy(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
	return arr
}
