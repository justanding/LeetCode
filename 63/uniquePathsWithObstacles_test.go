package uniquePathsWithObstacles_test

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Log(uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}))
	t.Log(uniquePathsWithObstacles([][]int{{0}}))
	t.Log(uniquePathsWithObstacles([][]int{{1}}))
}

/*
 * 在dp[m][n] = dp[m-1][n]+dp[m][n-1]的基础上加上障碍判断
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	n := len(obstacleGrid[0])

	obstacleGrid[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			obstacleGrid[i][0] = 0
		} else {
			obstacleGrid[i][0] = obstacleGrid[i-1][0]
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			obstacleGrid[0][j] = 0
		} else {
			obstacleGrid[0][j] = obstacleGrid[0][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[m-1][n-1]
}
