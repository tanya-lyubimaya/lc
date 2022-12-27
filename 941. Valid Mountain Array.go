package lc

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var peakPassed bool
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] && i > 1 {
			peakPassed = true
		}
		if !peakPassed && arr[i] <= arr[i-1] {
			return false
		}
		if peakPassed && arr[i] >= arr[i-1] {
			return false
		}
	}
	if !peakPassed {
		return false
	}
	return true
}

func validMountainArrayClassic(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var i int
	// walk up
	for i < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}
	// peak can't be first or last
	if i == 0 || i == len(arr)-1 {
		return false
	}
	// walk down
	for i < len(arr)-1 && arr[i] > arr[i+1] {
		i++
	}

	return i == len(arr)-1
}
