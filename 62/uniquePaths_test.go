package uniquePaths_test

import "testing"

func TestUniquePaths(t *testing.T) {
	t.Log(uniquePaths(1, 1))
	t.Log(uniquePaths(7, 3))
	t.Log(uniquePaths(20, 20))
}

/*
 * 采用动态规划
 * 边角部分都只有一种路径 path[m][0] = path[0][n] = 1
 * path[m][n] = path[m-1][n] + path[m][n-1]
 */
func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}
	for j := 0; j < n; j++ {
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}
