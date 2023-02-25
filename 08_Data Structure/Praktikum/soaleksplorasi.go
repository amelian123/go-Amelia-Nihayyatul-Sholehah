package main

import (
	"fmt"
	"math"
)

func diagonalDifference(matrix [][]int) int {
	n := len(matrix)
	leftDiag := 0
	rightDiag := 0
	for i := 0; i < n; i++ {
		leftDiag += matrix[i][i]
		rightDiag += matrix[i][n-i-1]
	}
	return int(math.Abs(float64(leftDiag - rightDiag)))
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	diff := diagonalDifference(matrix)
	fmt.Println(diff)
}
