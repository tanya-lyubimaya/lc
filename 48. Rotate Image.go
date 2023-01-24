package lc

func rotateWithMap(matrix [][]int) {
	numsWithCoordinates := make(map[[2]int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			numsWithCoordinates[[2]int{i, j}] = matrix[i][j]
		}
	}
	i := 0
	j1 := len(matrix[0]) - 1
	for i < len(matrix) && j1 >= 0 {
		j := 0
		i1 := 0
		for j < len(matrix[0]) && i1 < len(matrix) {
			matrix[i1][j1] = numsWithCoordinates[[2]int{i, j}]
			j++
			i1++
		}
		i++
		j1--
	}
}

func rotateWithNewMatrix(matrix [][]int) {
	oldMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			oldMatrix[i] = append(oldMatrix[i], matrix[i][j])
		}
	}
	i := 0
	j1 := len(matrix[0]) - 1
	for i < len(matrix) && j1 >= 0 {
		j := 0
		i1 := 0
		for j < len(matrix[0]) && i1 < len(matrix) {
			matrix[i1][j1] = oldMatrix[i][j]
			j++
			i1++
		}
		i++
		j1--
	}
}

func rotateGroups(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			temp := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = temp
		}
	}
}
