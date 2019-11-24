package minPathSum_test

import "testing"

func TestMinPathSum(t *testing.T) {
	t.Log(minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))

	t.Log(minPathSum([][]int{}))
}

/*
 * 动态规划
 * db[m][n] = grid[m][n] + min(grid[m-1][n],grid[m][n-1])
 */
func minPathSum(grid [][]int) int {
	m := len(grid)

	if m == 0 {
		return 0
	}

	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := grid[i][j-1]
			if min > grid[i-1][j] {
				min = grid[i-1][j]
			}
			grid[i][j] += min
		}
	}

	return grid[m-1][n-1]
}
