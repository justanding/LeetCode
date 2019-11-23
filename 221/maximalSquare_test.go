package maximalSquare_test

import (
	"testing"
)

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '1', '1', '0'},
	}
	t.Log(maximalSquare(matrix))

	matrix = [][]byte{
		[]byte{'0', '0', '0', '1'},
		[]byte{'1', '1', '0', '1'},
		[]byte{'1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1'},
		[]byte{'0', '1', '1', '1'},
	}
	t.Log(maximalSquare(matrix))

	t.Log(maximalSquare([][]byte{
		[]byte{'1'},
	}))
	t.Log(maximalSquare([][]byte{
		[]byte{'0'},
	}))
	t.Log(maximalSquare([][]byte{}))

}

/*
 * dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) + 1
 */
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	square := make([][]int, m)
	for i := 0; i < m; i++ {
		square[i] = make([]int, n)
	}

	maxSquare := 0
	for i := 0; i < m; i++ {
		square[i][0] = 0
		if matrix[i][0] == '1' {
			square[i][0] = 1
			maxSquare = 1
		}
	}

	for i := 0; i < n; i++ {
		square[0][i] = 0
		if matrix[0][i] == '1' {
			square[0][i] = 1
			maxSquare = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			square[i][j] = 0
			if matrix[i][j] == '1' {
				min := square[i-1][j-1]
				if min > square[i-1][j] {
					min = square[i-1][j]
				}
				if min > square[i][j-1] {
					min = square[i][j-1]
				}
				square[i][j] = min + 1
				if maxSquare < square[i][j] {
					maxSquare = square[i][j]
				}
			}
		}
	}

	return maxSquare * maxSquare
}
