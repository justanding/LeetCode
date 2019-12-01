package minimumTotal_test

import (
	_ "runtime/pprof"
	"testing"
)

func TestMinimumTotal(t *testing.T) {

	t.Log(minimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}))
	t.Log(minimumTotal([][]int{
		[]int{-1},
		[]int{3, 2},
		[]int{-3, 1, -1},
	}))
}

func BenchmarkMinimumTota(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumTotal([][]int{
			[]int{-1},
			[]int{3, 2},
			[]int{-3, 1, -1},
		})
	}
}

/*
 * https://leetcode-cn.com/problems/triangle/
 * dp[i] = min(dp[i-1],dp[i], dp[i+1])
 */
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}

	//初始化 m+1层中的数据，m+1个元素不使用，所以只初始化m个
	//m+1层无数据，所以都是0
	dp := make([]int, m)

	for i := m - 1; i >= 0; i-- {
		l := len(triangle[i])
		//第i层中，算出每个元素的值
		for j := 0; j < l; j++ {
			dp[j] += triangle[i][j]
		}

		//算出第i层每个元素j相邻的最小值
		for j := 0; j < l-1; j++ {
			if dp[j] > dp[j+1] {
				dp[j] = dp[j+1]
			}
		}
	}

	return dp[0]
}
